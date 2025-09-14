import { useState, useEffect } from 'react';
import { PropertySummary } from '@/types/api';
import { propertiesService } from '@/services/properties';

export interface UsePropertiesReturn {
  summary: PropertySummary | null;
  isLoading: boolean;
  error: string | null;
  refetch: () => Promise<void>;
}

export const useProperties = (): UsePropertiesReturn => {
  const [summary, setSummary] = useState<PropertySummary | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  const fetchSummary = async () => {
    try {
      setIsLoading(true);
      setError(null);
      const data = await propertiesService.getPropertySummary();
      setSummary(data);
    } catch (err) {
      setError(
        err instanceof Error ? err.message : 'Failed to fetch property data'
      );
      setSummary(null);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    fetchSummary();
  }, []);

  return {
    summary,
    isLoading,
    error,
    refetch: fetchSummary,
  };
};
