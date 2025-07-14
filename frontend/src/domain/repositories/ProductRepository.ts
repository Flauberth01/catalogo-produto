import { Product } from '../entities/Product';

export interface ProductFilter {
  name?: string;
  category?: string;
}

// Interface do repositório de produtos
export interface ProductRepository {
  getProducts(filters?: ProductFilter): Promise<Product[]>;
  getProduct(id: number): Promise<Product>;
  createProduct(product: Omit<Product, 'id'>): Promise<Product>;
  updateProduct(id: number, product: Partial<Product>): Promise<Product>;
  deleteProduct(id: number): Promise<void>;
} 