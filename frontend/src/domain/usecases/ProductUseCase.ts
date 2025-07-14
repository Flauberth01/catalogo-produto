import { Product } from '../entities/Product';
import { ProductRepository, ProductFilter } from '../repositories/ProductRepository';

// Casos de uso para produtos
export interface ProductUseCase {
  getProducts(filters?: ProductFilter): Promise<Product[]>;
  getProduct(id: number): Promise<Product>;
  createProduct(product: Omit<Product, 'id'>): Promise<Product>;
  updateProduct(id: number, product: Partial<Product>): Promise<Product>;
  deleteProduct(id: number): Promise<void>;
  searchProducts(query: string): Promise<Product[]>;
  filterProductsByCategory(category: string): Promise<Product[]>;
}

// Implementação dos casos de uso
export class ProductUseCaseImpl implements ProductUseCase {
  constructor(private productRepository: ProductRepository) {}

  async getProducts(filters?: ProductFilter): Promise<Product[]> {
    try {
      return await this.productRepository.getProducts(filters);
    } catch (error) {
      console.error('Erro ao buscar produtos:', error);
      throw new Error('Falha ao carregar produtos');
    }
  }

  async getProduct(id: number): Promise<Product> {
    try {
      return await this.productRepository.getProduct(id);
    } catch (error) {
      console.error('Erro ao buscar produto:', error);
      throw new Error('Produto não encontrado');
    }
  }

  async createProduct(product: Omit<Product, 'id'>): Promise<Product> {
    try {
      // Validações de negócio
      if (!product.name || product.name.trim().length === 0) {
        throw new Error('Nome do produto é obrigatório');
      }
      if (product.price <= 0) {
        throw new Error('Preço deve ser maior que zero');
      }

      return await this.productRepository.createProduct(product);
    } catch (error) {
      console.error('Erro ao criar produto:', error);
      throw error;
    }
  }

  async updateProduct(id: number, product: Partial<Product>): Promise<Product> {
    try {
      // Validações de negócio
      if (product.name !== undefined && product.name.trim().length === 0) {
        throw new Error('Nome do produto é obrigatório');
      }
      if (product.price !== undefined && product.price <= 0) {
        throw new Error('Preço deve ser maior que zero');
      }

      return await this.productRepository.updateProduct(id, product);
    } catch (error) {
      console.error('Erro ao atualizar produto:', error);
      throw error;
    }
  }

  async deleteProduct(id: number): Promise<void> {
    try {
      await this.productRepository.deleteProduct(id);
    } catch (error) {
      console.error('Erro ao deletar produto:', error);
      throw new Error('Falha ao deletar produto');
    }
  }

  async searchProducts(query: string): Promise<Product[]> {
    try {
      const products = await this.productRepository.getProducts();
      return products.filter(product => 
        product.name.toLowerCase().includes(query.toLowerCase()) ||
        product.description.toLowerCase().includes(query.toLowerCase())
      );
    } catch (error) {
      console.error('Erro ao buscar produtos:', error);
      throw new Error('Falha na busca de produtos');
    }
  }

  async filterProductsByCategory(category: string): Promise<Product[]> {
    try {
      return await this.productRepository.getProducts({ category });
    } catch (error) {
      console.error('Erro ao filtrar produtos por categoria:', error);
      throw new Error('Falha ao filtrar produtos');
    }
  }
} 