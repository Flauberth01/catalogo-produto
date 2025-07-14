import React from 'react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

// Configuração do QueryClient
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      // Configurações padrão para queries
      staleTime: 5 * 60 * 1000, // 5 minutos - dados considerados frescos
      gcTime: 10 * 60 * 1000, // 10 minutos - tempo no cache (antigo cacheTime)
      retry: 3, // Tentativas de retry
      refetchOnWindowFocus: false, // Não refetch ao focar na janela
      refetchOnReconnect: true, // Refetch ao reconectar
    },
    mutations: {
      // Configurações padrão para mutations
      retry: 1, // Apenas 1 tentativa para mutations
    },
  },
});

// Provider do React Query
interface QueryProviderProps {
  children: React.ReactNode;
}

export const QueryProvider: React.FC<QueryProviderProps> = ({ children }) => {
  return (
    <QueryClientProvider client={queryClient}>
      {children}
    </QueryClientProvider>
  );
};

// Exportar o queryClient para uso em outros lugares se necessário
export { queryClient }; 