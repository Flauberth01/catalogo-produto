import React, { useState, useMemo, useCallback } from 'react';
import Header from '../components/Header';
import Sidebar from '../components/Sidebar';
import ProductGrid from '../components/ProductGrid';
import { useAppQueries } from '../presentation/hooks/useAppQueries';

// Interface para filtros do frontend
interface FilterState {
  search: string;
  category: string;
}

const Index = () => {
  const [filters, setFilters] = useState<FilterState>({
    search: '',
    category: 'Todos'
  });
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false);

  const { products } = useAppQueries();
  const { data: productsData = [], isLoading: loading, error } = products;

  // Filtrar produtos localmente por busca e categoria
  const filteredProducts = useMemo(() => {
    let filtered = productsData;

    // Filtrar por categoria
    if (filters.category && filters.category !== 'Todos') {
      filtered = filtered.filter(product => product.category === filters.category);
    }

    // Filtrar por busca (nome e descrição)
    if (filters.search) {
      filtered = filtered.filter(product => {
        const matchesSearch = product.name.toLowerCase().includes(filters.search.toLowerCase()) ||
                             product.description.toLowerCase().includes(filters.search.toLowerCase());
        return matchesSearch;
      });
    }

    return filtered;
  }, [productsData, filters.search, filters.category]);

  const handleSearchChange = useCallback((search: string) => {
    setFilters(prev => ({ ...prev, search }));
  }, []);

  const handleCategoryChange = useCallback((category: string) => {
    setFilters(prev => ({ ...prev, category }));
  }, []);

  const handleMenuToggle = useCallback(() => {
    setIsMobileMenuOpen(!isMobileMenuOpen);
  }, [isMobileMenuOpen]);

  const closeMobileMenu = useCallback(() => {
    setIsMobileMenuOpen(false);
  }, []);

  return (
    <div className="min-h-screen bg-gray-50">
      <Header 
        onSearchChange={handleSearchChange}
        onMenuToggle={handleMenuToggle}
        isMobileMenuOpen={isMobileMenuOpen}
      />
      
      {/* Espaçamento para compensar o header fixo */}
      <div className="pt-16">
        <div className="flex">
          <Sidebar
            selectedCategory={filters.category}
            onCategoryChange={handleCategoryChange}
            isOpen={isMobileMenuOpen}
            onClose={closeMobileMenu}
          />
          
          <main className="flex-1 p-6 lg:pl-6">
            <div className="max-w-7xl mx-auto">
              {/* Breadcrumb / Status */}
              <div className="mb-6">
                <div className="flex items-center justify-between">
                  <div>
                    <h2 className="text-2xl font-bold text-gray-900">
                      {filters.category === 'Todos' ? 'Todos os Produtos' : filters.category}
                    </h2>
                    <p className="text-gray-600 mt-1">
                      {filteredProducts.length} {filteredProducts.length === 1 ? 'produto encontrado' : 'produtos encontrados'}
                      {filters.search && ` para "${filters.search}"`}
                    </p>
                  </div>
                </div>
              </div>

              {/* Error State */}
              {error && (
                <div className="text-center py-12">
                  <div className="text-red-500 text-lg mb-2">Erro ao carregar produtos</div>
                  <p className="text-gray-400">{error.message}</p>
                </div>
              )}

              {/* Products Grid */}
              <ProductGrid products={filteredProducts} loading={loading} />
            </div>
          </main>
        </div>
      </div>
    </div>
  );
};

export default Index;
