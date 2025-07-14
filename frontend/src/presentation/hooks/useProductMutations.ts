import { useMutation, useQueryClient } from '@tanstack/react-query';
import { ProductUseCase } from '../../domain/usecases/ProductUseCase';
import { Product } from '../../domain/entities/Product';
import { productKeys } from './useProductsQuery';

export const useCreateProduct = (productUseCase: ProductUseCase) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (product: Omit<Product, 'id'>) => productUseCase.createProduct(product),
    onSuccess: () => {
      // Invalidar cache de produtos para refetch
      queryClient.invalidateQueries({ queryKey: productKeys.lists() });
    },
    onError: (error) => {
      console.error('Erro ao criar produto:', error);
    },
  });
};

export const useUpdateProduct = (productUseCase: ProductUseCase) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, product }: { id: number; product: Partial<Product> }) =>
      productUseCase.updateProduct(id, product),
    onSuccess: (updatedProduct) => {
      // Invalidar cache de produtos
      queryClient.invalidateQueries({ queryKey: productKeys.lists() });
      // Atualizar produto específico no cache se existir
      queryClient.setQueryData(productKeys.detail(updatedProduct.id.toString()), updatedProduct);
    },
    onError: (error) => {
      console.error('Erro ao atualizar produto:', error);
    },
  });
};

export const useDeleteProduct = (productUseCase: ProductUseCase) => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number) => productUseCase.deleteProduct(id),
    onSuccess: (_, deletedId) => {
      // Invalidar cache de produtos
      queryClient.invalidateQueries({ queryKey: productKeys.lists() });
      // Remover produto específico do cache
      queryClient.removeQueries({ queryKey: productKeys.detail(deletedId.toString()) });
    },
    onError: (error) => {
      console.error('Erro ao deletar produto:', error);
    },
  });
}; 