import React, { useState } from 'react';
import ClientCard from '@/components/features/clients/ClientCard/ClientCard';
import { Client } from '@/components/features/clients/ClientCard/ClientCard.types';
import { Icon } from '@iconify/react';
import './Clients.css';

// Mock data - replace with actual API call in a real application
const mockClients: Client[] = [
  {
    id: '1',
    nome_do_cliente: 'João da Silva',
    telefone_do_cliente: '(11) 98765-4321',
    tipo_do_imovel: 'Apartamento',
    cidade_de_interesse: 'São Paulo',
    fluxo_do_cliente: 'Quente',
    preco_do_cliente: 850000,
    data_de_nascimento: '1985-05-15',
    renda_bruta_do_cliente: 15000,
    filhos_quantidade: 2,
    quantos_anos_de_carteira: 8,
    tem_outros_subsidios: true,
    possui_financiamento: false,
    fgts: true,
    local_habitado: 'Aluguel',
    financiado_quitado: 'Quitado',
    documentacao_dia: true,
    observacoes: 'Cliente interessado em apartamentos com 3 quartos e vaga de garagem.',
    conjuge: {
      data_de_nascimento_conjuge: '1987-09-22',
      renda_bruta_do_cliente_conjuge: 12000,
    },
  },
];

const Clients: React.FC = () => {
  const [clients] = useState<Client[]>(mockClients);
  const [searchTerm, setSearchTerm] = useState('');

  const handleEditClient = (client: Client) => {
    // Handle edit client logic
    console.log('Editing client:', client);
  };

  const handleCallClient = (phoneNumber: string) => {
    // Handle call client logic
    console.log('Calling:', phoneNumber);
    window.open(`tel:${phoneNumber}`, '_blank');
  };

  const handleViewSpouse = (spouseData: any) => {
    // Handle view spouse logic
    console.log('Viewing spouse:', spouseData);
  };

  const filteredClients = clients.filter(client =>
    client.nome_do_cliente.toLowerCase().includes(searchTerm.toLowerCase()) ||
    client.telefone_do_cliente.includes(searchTerm) ||
    client.cidade_de_interesse.toLowerCase().includes(searchTerm.toLowerCase())
  );

  return (
    <div className="clients-page">
      <div className="page-header">
        <h1>Clientes</h1>
        <p>Gerenciamento de clientes</p>
      </div>

      <div className="clients-actions">
        <div className="search-bar">
          <Icon icon="mdi:magnify" className="search-icon" />
          <input
            type="text"
            placeholder="Buscar clientes..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>
        <button className="btn btn-primary">
          <Icon icon="mdi:plus" /> Adicionar Cliente
        </button>
      </div>

      <div className="clients-grid">
        {filteredClients.length > 0 ? (
          filteredClients.map((client) => (
            <ClientCard
              key={client.id}
              client={client}
              onEdit={handleEditClient}
              onCall={handleCallClient}
              onViewSpouse={handleViewSpouse}
            />
          ))
        ) : (
          <div className="no-results">
            <Icon icon="mdi:account-question" className="no-results-icon" />
            <p>Nenhum cliente encontrado</p>
          </div>
        )}
      </div>
    </div>
  );
};

export default Clients;
