import { Property } from '@/types/api';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:3034';

export interface PropertyFilters {
  page?: number;
  limit?: number;
  search?: string;
  type?: string;
  city?: string;
  minPrice?: number;
  maxPrice?: number;
  status?: 'available' | 'sold' | 'all';
}

export interface PropertySummary {
  total: number;
  available: number;
  sold: number;
  averagePrice: number;
}

export interface PaginatedResponse<T> {
  data: T[];
  meta: {
    total: number;
    page: number;
    limit: number;
    totalPages: number;
    hasNextPage: boolean;
    hasPrevPage: boolean;
  };
}

export class PropertiesService {
  private async request<T>(
    endpoint: string,
    options?: RequestInit
  ): Promise<T> {
    const response = await fetch(`${API_BASE_URL}${endpoint}`, {
      ...options,
      headers: {
        'Content-Type': 'application/json',
        ...options?.headers,
      },
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    return response.json();
  }

  // Properties
  async getProperties(
    filters?: PropertyFilters
  ): Promise<PaginatedResponse<Property>> {
    const params = new URLSearchParams();
    if (filters) {
      Object.entries(filters).forEach(([key, value]) => {
        if (value !== undefined && value !== null) {
          params.append(key, value.toString());
        }
      });
    }

    return this.request<PaginatedResponse<Property>>(
      `/imoveis${params.toString() ? `?${params.toString()}` : ''}`
    );
  }

  async getPropertyById(id: number): Promise<Property> {
    return this.request<Property>(`/imoveis/${id}`);
  }

  // Property Summary
  async getPropertySummary(): Promise<PropertySummary> {
    const properties = await this.getProperties();
    const available = properties.data.filter(
      (p) => p.status === 'available'
    ).length;

    const total = properties.meta.total;
    const sold = total - available;
    const prices = properties.data.map((p) => p.price);
    const averagePrice =
      prices.length > 0
        ? Math.round(
            prices.reduce((sum, price) => sum + price, 0) / prices.length
          )
        : 0;

    return {
      total,
      available,
      sold,
      averagePrice,
    };
  }
}

export const propertiesService = new PropertiesService();
