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
  {
    id: '2',
    nome_do_cliente: 'Maria Santos',
    telefone_do_cliente: '(11) 99876-5432',
    tipo_do_imovel: 'Casa',
    cidade_de_interesse: 'Rio de Janeiro',
    fluxo_do_cliente: 'Morno',
    preco_do_cliente: 1200000,
    data_de_nascimento: '1990-03-10',
    renda_bruta_do_cliente: 18000,
    filhos_quantidade: 1,
    quantos_anos_de_carteira: 5,
    tem_outros_subsidios: false,
    possui_financiamento: true,
    fgts: false,
    local_habitado: 'Próprio',
    financiado_quitado: 'Em andamento',
    documentacao_dia: false,
    observacoes: 'Procurando casa com jardim grande para a família.',
  },
  {
    id: '3',
    nome_do_cliente: 'Carlos Oliveira',
    telefone_do_cliente: '(21) 98765-1234',
    tipo_do_imovel: 'Apartamento',
    cidade_de_interesse: 'Belo Horizonte',
    fluxo_do_cliente: 'Frio',
    preco_do_cliente: 600000,
    data_de_nascimento: '1982-11-25',
    renda_bruta_do_cliente: 12000,
    filhos_quantidade: 0,
    quantos_anos_de_carteira: 10,
    tem_outros_subsidios: true,
    possui_financiamento: false,
    fgts: true,
    local_habitado: 'Com os pais',
    financiado_quitado: 'Quitado',
    documentacao_dia: true,
    observacoes: 'Interessado em investimentos imobiliários.',
  },
  {
    id: '4',
    nome_do_cliente: 'Ana Pereira',
    telefone_do_cliente: '(11) 91234-5678',
    tipo_do_imovel: 'Casa',
    cidade_de_interesse: 'São Paulo',
    fluxo_do_cliente: 'Quente',
    preco_do_cliente: 950000,
    data_de_nascimento: '1988-07-18',
    renda_bruta_do_cliente: 16000,
    filhos_quantidade: 3,
    quantos_anos_de_carteira: 7,
    tem_outros_subsidios: false,
    possui_financiamento: false,
    fgts: true,
    local_habitado: 'Aluguel',
    financiado_quitado: 'Quitado',
    documentacao_dia: true,
    observacoes: 'Família crescendo, precisa de mais espaço.',
    conjuge: {
      data_de_nascimento_conjuge: '1986-12-05',
      renda_bruta_do_cliente_conjuge: 14000,
    },
  },
  {
    id: '5',
    nome_do_cliente: 'Pedro Costa',
    telefone_do_cliente: '(31) 99876-4321',
    tipo_do_imovel: 'Apartamento',
    cidade_de_interesse: 'Porto Alegre',
    fluxo_do_cliente: 'Morno',
    preco_do_cliente: 700000,
    data_de_nascimento: '1995-01-30',
    renda_bruta_do_cliente: 13000,
    filhos_quantidade: 0,
    quantos_anos_de_carteira: 3,
    tem_outros_subsidios: false,
    possui_financiamento: true,
    fgts: false,
    local_habitado: 'Aluguel',
    financiado_quitado: 'Em andamento',
    documentacao_dia: false,
    observacoes: 'Primeira compra, procurando apartamento moderno.',
  },
  {
    id: '6',
    nome_do_cliente: 'Juliana Lima',
    telefone_do_cliente: '(11) 98765-9876',
    tipo_do_imovel: 'Casa',
    cidade_de_interesse: 'São Paulo',
    fluxo_do_cliente: 'Quente',
    preco_do_cliente: 1100000,
    data_de_nascimento: '1983-09-14',
    renda_bruta_do_cliente: 20000,
    filhos_quantidade: 2,
    quantos_anos_de_carteira: 12,
    tem_outros_subsidios: true,
    possui_financiamento: false,
    fgts: true,
    local_habitado: 'Próprio',
    financiado_quitado: 'Quitado',
    documentacao_dia: true,
    observacoes: 'Quer mudar para uma casa maior com piscina.',
    conjuge: {
      data_de_nascimento_conjuge: '1981-04-20',
      renda_bruta_do_cliente_conjuge: 18000,
    },
  },
  {
    id: '7',
    nome_do_cliente: 'Roberto Alves',
    telefone_do_cliente: '(21) 91234-8765',
    tipo_do_imovel: 'Apartamento',
    cidade_de_interesse: 'Rio de Janeiro',
    fluxo_do_cliente: 'Frio',
    preco_do_cliente: 500000,
    data_de_nascimento: '1975-06-08',
    renda_bruta_do_cliente: 10000,
    filhos_quantidade: 1,
    quantos_anos_de_carteira: 15,
    tem_outros_subsidios: false,
    possui_financiamento: true,
    fgts: false,
    local_habitado: 'Próprio',
    financiado_quitado: 'Em andamento',
    documentacao_dia: false,
    observacoes: 'Aposentado procurando apartamento menor.',
  },
  {
    id: '8',
    nome_do_cliente: 'Fernanda Gomes',
    telefone_do_cliente: '(11) 99876-1234',
    tipo_do_imovel: 'Casa',
    cidade_de_interesse: 'Campinas',
    fluxo_do_cliente: 'Morno',
    preco_do_cliente: 800000,
    data_de_nascimento: '1992-02-12',
    renda_bruta_do_cliente: 14000,
    filhos_quantidade: 0,
    quantos_anos_de_carteira: 4,
    tem_outros_subsidios: true,
    possui_financiamento: false,
    fgts: true,
    local_habitado: 'Aluguel',
    financiado_quitado: 'Quitado',
    documentacao_dia: true,
    observacoes: 'Interessada em imóveis sustentáveis.',
  },
  {
    id: '9',
    nome_do_cliente: 'Marcos Silva',
    telefone_do_cliente: '(31) 98765-5678',
    tipo_do_imovel: 'Apartamento',
    cidade_de_interesse: 'Belo Horizonte',
    fluxo_do_cliente: 'Quente',
    preco_do_cliente: 650000,
    data_de_nascimento: '1987-08-27',
    renda_bruta_do_cliente: 11000,
    filhos_quantidade: 1,
    quantos_anos_de_carteira: 6,
    tem_outros_subsidios: false,
    possui_financiamento: false,
    fgts: true,
    local_habitado: 'Próprio',
    financiado_quitado: 'Quitado',
    documentacao_dia: true,
    observacoes: 'Procurando apartamento próximo ao trabalho.',
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
          <Icon icon="line-md:account-alert-loop" className="search-icon" />
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
