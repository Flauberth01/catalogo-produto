
import React from 'react';
import { ShoppingCart } from 'lucide-react';
import { Product } from '../domain/entities/Product';
import { useCart } from '../context/CartContext';

interface ProductCardProps {
  product: Product;
}

const ProductCard: React.FC<ProductCardProps> = ({ product }) => {
  const { addToCart } = useCart();

  const formatPrice = (price: number) => {
    return new Intl.NumberFormat('pt-BR', {
      style: 'currency',
      currency: 'BRL'
    }).format(price);
  };

  return (
    <div className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-300 group h-full flex flex-col">
      <div className="aspect-square overflow-hidden">
        <img
          src={product.image}
          alt={product.name}
          className="w-full h-full object-cover group-hover:scale-105 transition-transform duration-300"
          loading="lazy"
        />
      </div>
      
      <div className="p-4 flex flex-col flex-1">
        <div className="mb-2">
          <span className="inline-block bg-blue-100 text-blue-800 text-xs px-2 py-1 rounded-full">
            {product.category}
          </span>
        </div>
        
        <h3 className="text-lg font-semibold text-gray-900 mb-2 line-clamp-2">
          {product.name}
        </h3>
        
        <p className="text-gray-600 text-sm mb-3 line-clamp-3 flex-1">
          {product.description}
        </p>
        
        {/* Layout responsivo: stack vertical em telas pequenas, horizontal em telas maiores */}
        <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-3 mt-auto">
          <span className="text-xl sm:text-2xl font-bold text-green-600 flex-shrink-0">
            {formatPrice(product.price)}
          </span>
          
          <button
            onClick={() => addToCart(product)}
            className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg flex items-center justify-center gap-2 transition-colors duration-200 w-full sm:w-auto flex-shrink-0"
          >
            <ShoppingCart size={18} />
            <span>Adicionar</span>
          </button>
        </div>
      </div>
    </div>
  );
};

export default ProductCard;
