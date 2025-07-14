import { useProductsQuery } from './useProductsQuery';
import { useCategoriesQuery } from './useCategoriesQuery';
import { useCreateProduct, useUpdateProduct, useDeleteProduct } from './useProductMutations';
import { container } from '../../infrastructure/di/Container';

// Hook combinado para facilitar o uso dos hooks do React Query
export const useAppQueries = () => {
  const productUseCase = container.getProductUseCase();
  const categoryUseCase = container.getCategoryUseCase();

  // Queries
  const productsQuery = useProductsQuery(productUseCase);
  const categoriesQuery = useCategoriesQuery(categoryUseCase);

  // Mutations
  const createProductMutation = useCreateProduct(productUseCase);
  const updateProductMutation = useUpdateProduct(productUseCase);
  const deleteProductMutation = useDeleteProduct(productUseCase);

  return {
    // Queries
    products: productsQuery,
    categories: categoriesQuery,
    
    // Mutations
    createProduct: createProductMutation,
    updateProduct: updateProductMutation,
    deleteProduct: deleteProductMutation,
  };
}; 