export interface Property {
  id: number;
  title: string;
  description: string;
  price: number;
  type: string;
  status: 'available' | 'sold';
  city: string;
  neighborhood: string;
  bedrooms: number;
  bathrooms: number;
  area: number;
  images: string[];
  createdAt: string;
  updatedAt: string;
}

export interface PropertySummary {
  total: number;
  available: number;
  sold: number;
  averagePrice: number;
}
