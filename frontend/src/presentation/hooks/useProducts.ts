import { useState, useEffect, useCallback } from 'react';
import { Product } from '../../domain/entities/Product';
import { ProductFilter } from '../../domain/repositories/ProductRepository';
import { ProductUseCase } from '../../domain/usecases/ProductUseCase';

// Hook para gerenciar produtos
export const useProducts = (productUseCase: ProductUseCase) => {
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchProducts = useCallback(async (filters?: ProductFilter) => {
    try {
      setLoading(true);
      setError(null);
      const data = await productUseCase.getProducts(filters);
      setProducts(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Erro ao carregar produtos');
      console.error('Erro ao buscar produtos:', err);
    } finally {
      setLoading(false);
    }
  }, [productUseCase]);

  const searchProducts = useCallback(async (query: string) => {
    try {
      setLoading(true);
      setError(null);
      const data = await productUseCase.searchProducts(query);
      setProducts(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Erro na busca');
      console.error('Erro ao buscar produtos:', err);
    } finally {
      setLoading(false);
    }
  }, [productUseCase]);

  const filterByCategory = useCallback(async (category: string) => {
    try {
      setLoading(true);
      setError(null);
      const data = await productUseCase.filterProductsByCategory(category);
      setProducts(data);
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Erro ao filtrar produtos');
      console.error('Erro ao filtrar produtos:', err);
    } finally {
      setLoading(false);
    }
  }, [productUseCase]);

  const refetch = useCallback((filters?: ProductFilter) => {
    fetchProducts(filters);
  }, [fetchProducts]);

  return {
    products,
    loading,
    error,
    fetchProducts,
    searchProducts,
    filterByCategory,
    refetch,
  };
}; 