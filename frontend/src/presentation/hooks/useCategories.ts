import { useState, useEffect, useCallback } from 'react';
import { CategoryUseCase } from '../../domain/usecases/CategoryUseCase';

// Hook para gerenciar categorias
export const useCategories = (categoryUseCase: CategoryUseCase) => {
  const [categories, setCategories] = useState<string[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  const fetchCategories = useCallback(async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await categoryUseCase.getCategories();
      setCategories(['Todos', ...data]); // Adicionar "Todos" como primeira opção
    } catch (err) {
      setError(err instanceof Error ? err.message : 'Erro ao carregar categorias');
      console.error('Erro ao buscar categorias:', err);
    } finally {
      setLoading(false);
    }
  }, [categoryUseCase]);

  const refetch = useCallback(() => {
    fetchCategories();
  }, [fetchCategories]);

  return {
    categories,
    loading,
    error,
    fetchCategories,
    refetch,
  };
}; 