import React from 'react';
import { Meta, StoryObj } from '@storybook/react';
import ClientCard, { Client } from './ClientCard';

export default {
  title: 'Features/Clients/ClientCard',
  component: ClientCard,
  tags: ['autodocs'],
  argTypes: {
    client: { control: 'object' },
    onEdit: { action: 'edit' },
    onCall: { action: 'call' },
    onViewMedia: { action: 'viewMedia' },
    onViewSpouse: { action: 'viewSpouse' },
  },
} as Meta<typeof ClientCard>;

// Sample client data
const sampleClient: Client = {
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
  media: [
    'https://example.com/photo1.jpg',
    'https://example.com/photo2.jpg',
  ],
};

type Story = StoryObj<typeof ClientCard>;

export const Default: Story = {
  args: {
    client: sampleClient,
  },
};

export const WithoutSpouse: Story = {
  args: {
    client: {
      ...sampleClient,
      conjuge: undefined,
    },
  },
};

export const WithoutMedia: Story = {
  args: {
    client: {
      ...sampleClient,
      media: [],
    },
  },
};

export const MinimalInfo: Story = {
  args: {
    client: {
      id: '2',
      nome_do_cliente: 'Maria Oliveira',
      telefone_do_cliente: '(11) 91234-5678',
      tipo_do_imovel: 'Casa',
      cidade_de_interesse: 'Campinas',
      fluxo_do_cliente: 'Morno',
      preco_do_cliente: 1200000,
    },
  },
};

export const ColdLead: Story => ({
  args: {
    client: {
      ...sampleClient,
      id: '3',
      nome_do_cliente: 'Carlos Souza',
      fluxo_do_cliente: 'Frio',
      preco_do_cliente: 2000000,
      cidade_de_interesse: 'São José dos Campos',
      observacoes: 'Cliente em fase inicial de contato. Interessado em casas de luxo.',
    },
  },
});

export const WithLongText: Story = {
  args: {
    client: {
      ...sampleClient,
      id: '4',
      nome_do_cliente: 'Ana Carolina de Oliveira Santos Pereira',
      observacoes: 'Este é um exemplo de um texto muito longo que pode ser incluído no campo de observações. O componente deve lidar bem com textos longos, quebrando as linhas corretamente e garantindo que o layout permaneça consistente e legível, sem quebrar o design do card.',
    },
  },
};

// Function to create a grid of client cards
export const GridView: Story = {
  render: () => (
    <div style={{
      display: 'grid',
      gridTemplateColumns: 'repeat(auto-fill, minmax(350px, 1fr))',
      gap: '20px',
      padding: '20px',
    }}>
      <ClientCard client={sampleClient} />
      <ClientCard client={{
        ...sampleClient,
        id: '2',
        nome_do_cliente: 'Maria Oliveira',
        fluxo_do_cliente: 'Morno',
        preco_do_cliente: 1200000,
        cidade_de_interesse: 'Campinas',
      }} />
      <ClientCard client={{
        ...sampleClient,
        id: '3',
        nome_do_cliente: 'Carlos Souza',
        fluxo_do_cliente: 'Frio',
        preco_do_cliente: 2000000,
        cidade_de_interesse: 'São José dos Campos',
      }} />
    </div>
  ),
};

// Add a decorator to show how it looks in a container
export const InContainer: Story = {
  args: {
    client: sampleClient,
  },
  decorators: [
    (Story) => (
      <div style={{ maxWidth: '800px', margin: '0 auto', padding: '20px' }}>
        <h2>Lista de Clientes</h2>
        <div style={{ display: 'flex', flexDirection: 'column', gap: '16px' }}>
          <Story />
          <ClientCard client={{
            ...sampleClient,
            id: '2',
            nome_do_cliente: 'Maria Oliveira',
            fluxo_do_cliente: 'Morno',
          }} />
        </div>
      </div>
    ),
  ],
};
