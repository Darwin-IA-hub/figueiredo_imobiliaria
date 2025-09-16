import { ReactNode } from 'react';

export interface Client {
  id: string;
  nome_do_cliente: string;
  telefone_do_cliente: string;
  tipo_do_imovel: string;
  cidade_de_interesse: string;
  fluxo_do_cliente: string;
  preco_do_cliente: number;
  observacoes?: string;
  data_de_nascimento?: string;
  renda_bruta_do_cliente?: number;
  filhos_quantidade?: number;
  quantos_anos_de_carteira?: number;
  tem_outros_subsidios?: boolean;
  possui_financiamento?: boolean;
  fgts?: boolean;
  local_habitado?: string;
  financiado_quitado?: string;
  documentacao_dia?: boolean;
  conjuge?: {
    data_de_nascimento_conjuge?: string;
    renda_bruta_do_cliente_conjuge?: number;
  };
  media?: string[];
}

export interface ClientCardProps {
  client: Client;
  className?: string;
  onEdit?: (client: Client) => void;
  onCall?: (phoneNumber: string) => void;
  onViewMedia?: (media: string[]) => void;
  onViewSpouse?: (spouseData: Client['conjuge']) => void;
}

export interface InfoItemProps {
  icon: ReactNode;
  label: string;
  value?: string | number | boolean | null;
  className?: string;
  hideIfEmpty?: boolean;
  formatValue?: (value: any) => ReactNode | string;
}

export interface ExpandableSectionProps {
  isOpen: boolean;
  children: ReactNode;
  className?: string;
}
