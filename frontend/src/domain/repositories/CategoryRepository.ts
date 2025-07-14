import { Category } from '../entities/Product';

// Interface do repositório de categorias
export interface CategoryRepository {
  getCategories(): Promise<string[]>;
  getCategory(id: number): Promise<Category>;
  getCategoryByName(name: string): Promise<Category | null>;
  createCategory(name: string): Promise<Category>;
  updateCategory(id: number, name: string): Promise<Category>;
  deleteCategory(id: number): Promise<void>;
} 