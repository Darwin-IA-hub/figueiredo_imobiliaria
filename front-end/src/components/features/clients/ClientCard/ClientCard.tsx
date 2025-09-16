import React, { useState, useMemo } from 'react';
// Using native Date methods instead of date-fns to avoid dependency issues
const formatDate = (dateString?: string) => {
  if (!dateString) return 'Não informado';
  try {
    const date = new Date(dateString);
    return date.toLocaleDateString('pt-BR');
  } catch (error) {
    return 'Data inválida';
  }
};

// Date formatting is handled by native toLocaleDateString
import { Icon } from '@iconify/react';
import { ClientCardProps, InfoItemProps, ExpandableSectionProps } from './ClientCard.types';
import './ClientCard.css';

// Helper component for info items
const InfoItem: React.FC<InfoItemProps> = ({
  icon,
  label,
  value,
  className = '',
  hideIfEmpty = true,
  formatValue = (val) => String(val ?? '')
}) => {
  if (hideIfEmpty && (value === undefined || value === null || value === '')) {
    return null;
  }

  return (
    <div className={`info-item ${className}`}>
      <span className="info-icon">{icon}</span>
      <div className="info-content">
        <span className="info-label">{label}:</span>
        <span className="info-value">{formatValue(value)}</span>
      </div>
    </div>
  );
};

// Expandable section component
const ExpandableSection: React.FC<ExpandableSectionProps> = ({
  isOpen,
  children,
  className = ''
}) => {
  return (
    <div className={`expandable-section ${isOpen ? 'expanded' : ''} ${className}`}>
      <div className="expandable-content">{children}</div>
    </div>
  );
};

// Format currency
const formatCurrency = (value?: number) => {
  if (value === undefined || value === null) return 'Não informado';
  return new Intl.NumberFormat('pt-BR', {
    style: 'currency',
    currency: 'BRL',
    minimumFractionDigits: 2,
    maximumFractionDigits: 2
  }).format(value);
};

// Format date is now defined at the top

// Format boolean to Sim/Não
const formatBoolean = (value?: boolean) => (value ? 'Sim' : 'Não');

const ClientCard: React.FC<ClientCardProps> = ({
  client,
  className = '',
  onEdit,
  onCall,
  onViewMedia,
  onViewSpouse
}) => {
  const [isExpanded, setIsExpanded] = useState(false);

  // Toggle card expansion
  const toggleExpand = () => {
    setIsExpanded(!isExpanded);
  };

  // Handle phone number click
  const handlePhoneClick = (e: React.MouseEvent, phoneNumber: string) => {
    e.stopPropagation();
    if (onCall) {
      onCall(phoneNumber);
    } else {
      window.location.href = `tel:${phoneNumber}`;
    }
  };

  // Handle edit button click
  const handleEditClick = (e: React.MouseEvent) => {
    e.stopPropagation();
    onEdit?.(client);
  };

  // Handle media button click
  const handleMediaClick = (e: React.MouseEvent) => {
    e.stopPropagation();
    onViewMedia?.(client.media || []);
  };

  // Handle spouse button click
  const handleSpouseClick = (e: React.MouseEvent) => {
    e.stopPropagation();
    onViewSpouse?.(client.conjuge);
  };

  // Get status color based on client status
  const statusColor = useMemo(() => {
    switch (client.fluxo_do_cliente?.toLowerCase()) {
      case 'quente':
        return 'status-hot';
      case 'morno':
        return 'status-warm';
      case 'frio':
        return 'status-cold';
      default:
        return 'status-default';
    }
  }, [client.fluxo_do_cliente]);

  return (
    <div 
      className={`client-card ${className} ${isExpanded ? 'expanded' : ''}`}
      onClick={toggleExpand}
    >
      {/* Card Header */}
      <div className="card-header">
        <div className="client-name-status">
          <h3 className="client-name">{client.nome_do_cliente}</h3>
          <span className={`status-badge ${statusColor}`}>
            {client.fluxo_do_cliente || 'Sem status'}
          </span>
        </div>
        
        <div className="card-actions">
          {client.conjuge && (
            <button 
              className="icon-button" 
              onClick={handleSpouseClick}
              aria-label="Ver cônjuge"
              title="Ver cônjuge"
            >
              <Icon icon="mdi:account-group" />
            </button>
          )}
          
          {client.media && client.media.length > 0 && (
            <button 
              className="icon-button" 
              onClick={handleMediaClick}
              aria-label="Ver mídia"
              title="Ver mídia"
            >
              <Icon icon="mdi:image-multiple" />
            </button>
          )}
          
          <button 
            className="icon-button" 
            onClick={handleEditClick}
            aria-label="Editar"
            title="Editar"
          >
            <Icon icon="mdi:pencil" />
          </button>
          
          <button 
            className={`expand-button ${isExpanded ? 'expanded' : ''}`}
            aria-label={isExpanded ? 'Recolher' : 'Expandir'}
            title={isExpanded ? 'Recolher' : 'Expandir'}
          >
            <Icon icon={isExpanded ? "mdi:chevron-up" : "mdi:chevron-down"} />
          </button>
        </div>
      </div>

      {/* Primary Info */}
      <div className="primary-info">
        <InfoItem
          icon={<Icon icon="mdi:phone" />}
          label="Telefone"
          value={client.telefone_do_cliente}
          formatValue={(value) => (
            <a 
              href={`tel:${value}`} 
              className="phone-link"
              onClick={(e) => handlePhoneClick(e, value as string)}
            >
              {value}
            </a>
          )}
        />
        
        <InfoItem
          icon={<Icon icon="mdi:city" />}
          label="Cidade de interesse"
          value={client.cidade_de_interesse}
        />
        
        <InfoItem
          icon={<Icon icon="mdi:home" />}
          label="Tipo de imóvel"
          value={client.tipo_do_imovel}
        />
        
        <InfoItem
          icon={<Icon icon="mdi:cash" />}
          label="Orçamento"
          value={client.preco_do_cliente}
          formatValue={formatCurrency}
        />
      </div>

      {/* Expandable Section */}
      <ExpandableSection isOpen={isExpanded}>
        <div className="secondary-info">
          <h4>Informações Adicionais</h4>
          
          <div className="info-grid">
            <InfoItem
              icon={<Icon icon="mdi:cake" />}
              label="Data de Nascimento"
              value={client.data_de_nascimento}
              formatValue={formatDate}
            />
            
            <InfoItem
              icon={<Icon icon="mdi:cash-multiple" />}
              label="Renda Bruta"
              value={client.renda_bruta_do_cliente}
              formatValue={formatCurrency}
            />
            
            <InfoItem
              icon={<Icon icon="mdi:human-male-child" />}
              label="Filhos"
              value={client.filhos_quantidade}
              formatValue={(value) => value ? `${value} ${value === 1 ? 'filho' : 'filhos'}` : 'Nenhum'}
            />
            
            <InfoItem
              icon={<Icon icon="mdi:card-account-details" />}
              label="Anos de Carteira"
              value={client.quantos_anos_de_carteira}
            />
            
            <InfoItem
              icon={<Icon icon="mdi:home-account" />}
              label="Local Habitado"
              value={client.local_habitado}
            />
            
            <InfoItem
              icon={<Icon icon="mdi:home-currency-usd" />}
              label="Financiado/Quitado"
              value={client.financiado_quitado}
            />
            
            <InfoItem
              icon={<Icon icon="mdi:file-document-check" />}
              label="Documentação em dia"
              value={client.documentacao_dia}
              formatValue={formatBoolean}
            />
            
            <InfoItem
              icon={<Icon icon="mdi:hand-coin" />}
              label="Usa FGTS"
              value={client.fgts}
              formatValue={formatBoolean}
            />
            
            <InfoItem
              icon={<Icon icon="mdi:home-plus" />}
              label="Possui Financiamento"
              value={client.possui_financiamento}
              formatValue={formatBoolean}
            />
            
            <InfoItem
              icon={<Icon icon="mdi:hand-coin-outline" />}
              label="Outros Subsídios"
              value={client.tem_outros_subsidios}
              formatValue={formatBoolean}
            />
          </div>
          
          {client.observacoes && (
            <div className="notes-section">
              <h5>Observações</h5>
              <p className="notes-content">{client.observacoes}</p>
            </div>
          )}
        </div>
      </ExpandableSection>
    </div>
  );
};

export default ClientCard;
