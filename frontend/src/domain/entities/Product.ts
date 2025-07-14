// Entidade de domínio para Product
export interface Product {
  id: number;
  name: string;
  price: number;
  image: string;
  category: string;
  description: string;
}

// Entidade de domínio para CartItem
export interface CartItem extends Product {
  quantity: number;
}

// Entidade de domínio para Category
export interface Category {
  id: number;
  name: string;
} 