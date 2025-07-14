import { Category } from '../entities/Product';
import { CategoryRepository } from '../repositories/CategoryRepository';

// Casos de uso para categorias
export interface CategoryUseCase {
  getCategories(): Promise<string[]>;
  getCategory(id: number): Promise<Category>;
  getCategoryByName(name: string): Promise<Category | null>;
  createCategory(name: string): Promise<Category>;
  updateCategory(id: number, name: string): Promise<Category>;
  deleteCategory(id: number): Promise<void>;
}

// Implementação dos casos de uso
export class CategoryUseCaseImpl implements CategoryUseCase {
  constructor(private categoryRepository: CategoryRepository) {}

  async getCategories(): Promise<string[]> {
    try {
      return await this.categoryRepository.getCategories();
    } catch (error) {
      console.error('Erro ao buscar categorias:', error);
      throw new Error('Falha ao carregar categorias');
    }
  }

  async getCategory(id: number): Promise<Category> {
    try {
      return await this.categoryRepository.getCategory(id);
    } catch (error) {
      console.error('Erro ao buscar categoria:', error);
      throw new Error('Categoria não encontrada');
    }
  }

  async getCategoryByName(name: string): Promise<Category | null> {
    try {
      return await this.categoryRepository.getCategoryByName(name);
    } catch (error) {
      console.error('Erro ao buscar categoria por nome:', error);
      throw new Error('Falha ao buscar categoria');
    }
  }

  async createCategory(name: string): Promise<Category> {
    try {
      // Validações de negócio
      if (!name || name.trim().length === 0) {
        throw new Error('Nome da categoria é obrigatório');
      }
      if (name.trim().length < 2) {
        throw new Error('Nome da categoria deve ter pelo menos 2 caracteres');
      }

      return await this.categoryRepository.createCategory(name.trim());
    } catch (error) {
      console.error('Erro ao criar categoria:', error);
      throw error;
    }
  }

  async updateCategory(id: number, name: string): Promise<Category> {
    try {
      // Validações de negócio
      if (!name || name.trim().length === 0) {
        throw new Error('Nome da categoria é obrigatório');
      }
      if (name.trim().length < 2) {
        throw new Error('Nome da categoria deve ter pelo menos 2 caracteres');
      }

      return await this.categoryRepository.updateCategory(id, name.trim());
    } catch (error) {
      console.error('Erro ao atualizar categoria:', error);
      throw error;
    }
  }

  async deleteCategory(id: number): Promise<void> {
    try {
      await this.categoryRepository.deleteCategory(id);
    } catch (error) {
      console.error('Erro ao deletar categoria:', error);
      throw new Error('Falha ao deletar categoria');
    }
  }
} 