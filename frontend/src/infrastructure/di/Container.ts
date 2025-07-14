import { ProductRepository } from '../../domain/repositories/ProductRepository';
import { CategoryRepository } from '../../domain/repositories/CategoryRepository';
import { ProductUseCase } from '../../domain/usecases/ProductUseCase';
import { CategoryUseCase } from '../../domain/usecases/CategoryUseCase';
import { ProductRepositoryImpl } from '../repositories/ProductRepositoryImpl';
import { CategoryRepositoryImpl } from '../repositories/CategoryRepositoryImpl';
import { ProductUseCaseImpl } from '../../domain/usecases/ProductUseCase';
import { CategoryUseCaseImpl } from '../../domain/usecases/CategoryUseCase';

// Container de Injeção de Dependências
export class Container {
  private static instance: Container;
  private productRepository: ProductRepository;
  private categoryRepository: CategoryRepository;
  private productUseCase: ProductUseCase;
  private categoryUseCase: CategoryUseCase;

  private constructor() {
    // Inicializar repositórios
    this.productRepository = new ProductRepositoryImpl();
    this.categoryRepository = new CategoryRepositoryImpl();

    // Inicializar casos de uso
    this.productUseCase = new ProductUseCaseImpl(this.productRepository);
    this.categoryUseCase = new CategoryUseCaseImpl(this.categoryRepository);
  }

  // Singleton pattern
  public static getInstance(): Container {
    if (!Container.instance) {
      Container.instance = new Container();
    }
    return Container.instance;
  }

  // Getters para as dependências
  public getProductUseCase(): ProductUseCase {
    return this.productUseCase;
  }

  public getCategoryUseCase(): CategoryUseCase {
    return this.categoryUseCase;
  }

  public getProductRepository(): ProductRepository {
    return this.productRepository;
  }

  public getCategoryRepository(): CategoryRepository {
    return this.categoryRepository;
  }
}

// Instância global do container
export const container = Container.getInstance(); 