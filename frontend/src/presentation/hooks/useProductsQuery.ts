import { useQuery } from '@tanstack/react-query';
import { ProductUseCase } from '../../domain/usecases/ProductUseCase';

// Chaves para o React Query
export const productKeys = {
  all: ['products'] as const,
  lists: () => [...productKeys.all, 'list'] as const,
  list: (filters: string) => [...productKeys.lists(), { filters }] as const,
  details: () => [...productKeys.all, 'detail'] as const,
  detail: (id: string) => [...productKeys.details(), id] as const,
};

export const useProductsQuery = (productUseCase: ProductUseCase) => {
  return useQuery({
    queryKey: productKeys.lists(),
    queryFn: () => productUseCase.getProducts(),
    staleTime: 5 * 60 * 1000, // 5 minutos
    gcTime: 10 * 60 * 1000, // 10 minutos
  });
}; 