import { Category } from '../../domain/entities/Product';
import { CategoryRepository } from '../../domain/repositories/CategoryRepository';
import { apiClient } from '../api/ApiClient';

// DTOs para comunicação com a API
interface ApiCategory {
  id: number;
  name: string;
  created_at: string;
  updated_at: string;
}

interface ApiResponse<T> {
  data: T;
  total?: number;
}

// Implementação do repositório de categorias
export class CategoryRepositoryImpl implements CategoryRepository {
  // Buscar todas as categorias
  async getCategories(): Promise<string[]> {
    try {
      const response = await apiClient.get<ApiResponse<ApiCategory[]>>('/categories');
      return response.data.map(category => category.name);
    } catch (error) {
      console.error('Erro ao buscar categorias:', error);
      throw error;
    }
  }

  // Buscar categoria por ID
  async getCategory(id: number): Promise<Category> {
    try {
      const response = await apiClient.get<ApiResponse<ApiCategory>>(`/categories/${id}`);
      return {
        id: response.data.id,
        name: response.data.name,
      };
    } catch (error) {
      console.error('Erro ao buscar categoria:', error);
      throw error;
    }
  }

  // Buscar categoria por nome
  async getCategoryByName(name: string): Promise<Category | null> {
    try {
      const categories = await apiClient.get<ApiResponse<ApiCategory[]>>('/categories');
      const category = categories.data.find(cat => cat.name === name);
      return category ? { id: category.id, name: category.name } : null;
    } catch (error) {
      console.error('Erro ao buscar categoria por nome:', error);
      throw error;
    }
  }

  // Criar nova categoria
  async createCategory(name: string): Promise<Category> {
    try {
      const response = await apiClient.post<ApiResponse<ApiCategory>>('/categories', { name });
      return {
        id: response.data.id,
        name: response.data.name,
      };
    } catch (error) {
      console.error('Erro ao criar categoria:', error);
      throw error;
    }
  }

  // Atualizar categoria
  async updateCategory(id: number, name: string): Promise<Category> {
    try {
      const response = await apiClient.put<ApiResponse<ApiCategory>>(`/categories/${id}`, { name });
      return {
        id: response.data.id,
        name: response.data.name,
      };
    } catch (error) {
      console.error('Erro ao atualizar categoria:', error);
      throw error;
    }
  }

  // Deletar categoria
  async deleteCategory(id: number): Promise<void> {
    try {
      await apiClient.delete(`/categories/${id}`);
    } catch (error) {
      console.error('Erro ao deletar categoria:', error);
      throw error;
    }
  }
} 