import React from 'react';
import { useAppQueries } from '../presentation/hooks/useAppQueries';

interface SidebarProps {
  selectedCategory: string;
  onCategoryChange: (category: string) => void;
  isOpen: boolean;
  onClose: () => void;
}

const Sidebar: React.FC<SidebarProps> = ({ 
  selectedCategory, 
  onCategoryChange, 
  isOpen, 
  onClose 
}) => {
  const { categories } = useAppQueries();
  const { data: categoriesData = [], isLoading: loading, error } = categories;

  if (loading) {
    return (
      <aside className="fixed lg:static inset-y-0 left-0 z-50 w-64 bg-white shadow-lg lg:shadow-none lg:border-r border-gray-200">
        <div className="p-6 pt-20 lg:pt-6">
          <h2 className="text-lg font-semibold text-gray-900 mb-4">Categorias</h2>
          <div className="space-y-2">
            {Array.from({ length: 5 }).map((_, index) => (
              <div key={index} className="h-8 bg-gray-200 rounded animate-pulse"></div>
            ))}
          </div>
        </div>
      </aside>
    );
  }

  if (error) {
    return (
      <aside className="fixed lg:static inset-y-0 left-0 z-50 w-64 bg-white shadow-lg lg:shadow-none lg:border-r border-gray-200">
        <div className="p-6 pt-20 lg:pt-6">
          <h2 className="text-lg font-semibold text-gray-900 mb-4">Categorias</h2>
          <div className="text-red-500 text-sm">Erro ao carregar categorias</div>
        </div>
      </aside>
    );
  }

  return (
    <>
      {/* Overlay para mobile */}
      {isOpen && (
        <div 
          className="fixed inset-0 bg-black bg-opacity-50 z-40 lg:hidden"
          onClick={onClose}
        />
      )}

      {/* Sidebar */}
      <aside 
        className={`
          fixed lg:static inset-y-0 left-0 z-50 w-64 bg-white shadow-lg transform transition-transform duration-300 ease-in-out
          ${isOpen ? 'translate-x-0' : '-translate-x-full'} 
          lg:translate-x-0 lg:shadow-none lg:border-r border-gray-200
        `}
      >
        <div className="p-6 pt-20 lg:pt-6">
          <h2 className="text-lg font-semibold text-gray-900 mb-4">Categorias</h2>
          
          <div className="space-y-2">
            {["Todos", ...categoriesData.filter(c => c !== "Todos")].map((category) => (
              <button
                key={category}
                onClick={() => {
                  onCategoryChange(category);
                  onClose(); // Fechar sidebar no mobile após selecionar
                }}
                className={`
                  w-full text-left px-3 py-2 rounded-lg transition-colors duration-200
                  ${selectedCategory === category
                    ? 'bg-blue-100 text-blue-700 font-medium'
                    : 'text-gray-700 hover:bg-gray-100'
                  }
                `}
              >
                {category}
              </button>
            ))}
          </div>
        </div>
      </aside>
    </>
  );
};

export default Sidebar;
