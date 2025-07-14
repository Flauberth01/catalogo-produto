import { useQuery } from '@tanstack/react-query';
import { CategoryUseCase } from '../../domain/usecases/CategoryUseCase';

// Chaves para o React Query
export const categoryKeys = {
  all: ['categories'] as const,
  lists: () => [...categoryKeys.all, 'list'] as const,
  list: (filters: string) => [...categoryKeys.lists(), { filters }] as const,
  details: () => [...categoryKeys.all, 'detail'] as const,
  detail: (id: string) => [...categoryKeys.details(), id] as const,
};

export const useCategoriesQuery = (categoryUseCase: CategoryUseCase) => {
  return useQuery({
    queryKey: categoryKeys.lists(),
    queryFn: () => categoryUseCase.getCategories(),
    staleTime: 10 * 60 * 1000, // 10 minutos - categorias mudam menos
    gcTime: 15 * 60 * 1000, // 15 minutos
  });
}; 