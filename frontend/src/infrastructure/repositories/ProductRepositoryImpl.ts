import { Product } from '../../domain/entities/Product';
import { ProductRepository, ProductFilter } from '../../domain/repositories/ProductRepository';
import { apiClient } from '../api/ApiClient';

// DTOs para comunicação com a API
interface ApiProduct {
  id: number;
  name: string;
  price: number;
  image: string;
  category_id: number;
  category: {
    id: number;
    name: string;
  };
  description: string;
  created_at: string;
  updated_at: string;
}

interface ApiResponse<T> {
  data: T;
  total?: number;
}

// Implementação do repositório de produtos
export class ProductRepositoryImpl implements ProductRepository {
  // Converter produto da API para o formato do domínio
  private mapApiProductToProduct(apiProduct: ApiProduct): Product {
    return {
      id: apiProduct.id,
      name: apiProduct.name,
      price: apiProduct.price,
      image: apiProduct.image,
      category: apiProduct.category.name,
      description: apiProduct.description,
    };
  }

  // Buscar todos os produtos com filtros opcionais
  async getProducts(filters?: ProductFilter): Promise<Product[]> {
    try {
      const params = new URLSearchParams();
      if (filters?.name) params.append('name', filters.name);
      if (filters?.category) params.append('category', filters.category);

      const response = await apiClient.get<ApiResponse<ApiProduct[]>>(`/products?${params.toString()}`);
      return response.data.map(this.mapApiProductToProduct);
    } catch (error) {
      console.error('Erro ao buscar produtos:', error);
      throw error;
    }
  }

  // Buscar produto por ID
  async getProduct(id: number): Promise<Product> {
    try {
      const response = await apiClient.get<ApiResponse<ApiProduct>>(`/products/${id}`);
      return this.mapApiProductToProduct(response.data);
    } catch (error) {
      console.error('Erro ao buscar produto:', error);
      throw error;
    }
  }

  // Criar novo produto
  async createProduct(product: Omit<Product, 'id'>): Promise<Product> {
    try {
      const response = await apiClient.post<ApiResponse<ApiProduct>>('/products', {
        name: product.name,
        price: product.price,
        image: product.image,
        description: product.description,
        category_id: 1, // TODO: Implementar busca de categoria por nome
      });
      return this.mapApiProductToProduct(response.data);
    } catch (error) {
      console.error('Erro ao criar produto:', error);
      throw error;
    }
  }

  // Atualizar produto
  async updateProduct(id: number, product: Partial<Product>): Promise<Product> {
    try {
      const response = await apiClient.put<ApiResponse<ApiProduct>>(`/products/${id}`, {
        name: product.name,
        price: product.price,
        image: product.image,
        description: product.description,
        category_id: 1, // TODO: Implementar busca de categoria por nome
      });
      return this.mapApiProductToProduct(response.data);
    } catch (error) {
      console.error('Erro ao atualizar produto:', error);
      throw error;
    }
  }

  // Deletar produto
  async deleteProduct(id: number): Promise<void> {
    try {
      await apiClient.delete(`/products/${id}`);
    } catch (error) {
      console.error('Erro ao deletar produto:', error);
      throw error;
    }
  }
} 