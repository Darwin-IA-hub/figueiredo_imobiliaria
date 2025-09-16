import React from 'react';
import './Footer.css';

const Footer: React.FC = () => {
  return (
    <footer className="footer">
      <div className="footer-content">
        <div className="footer-section">
          <h3>Horário de Atendimento</h3>
          <p>De segunda a Sexta-feira</p>
          <p>Horário das 9h às 17h</p>
          <p>Intervalo para almoço das 12h às 14h</p>
        </div>

        <div className="footer-section">
          <h3>Contato</h3>
          <p>Site: <a href="https://figueiredoimoveis.com.br" target="_blank" rel="noopener noreferrer">figueiredoimoveis.com.br</a></p>
          <p>E-mail: contato@figueiredoimoveis.com.br</p>
          <p>Consultores: (15) 98151-7070</p>
          <p>Fone da placa: (15) 98151-7070</p>
        </div>

        <div className="footer-section">
          <h3>Endereço</h3>
          <p>Av. Cel Firmo Vieira de Camargo, 825</p>
          <p>Centro - Tatuí/SP</p>
          <p>Desde março de 1996 - Há mais de 25 anos</p>
        </div>
      </div>

      <div className="footer-bottom">
        <p>© {new Date().getFullYear()} Figueiredo Imobiliária - Imóvel moeda forte</p>
        <p>Atuamos em vendas de lançamentos, ajudando famílias a realizarem o sonho de adquirir seu imóvel para morar e investir.</p>
      </div>
    </footer>
  );
};

export default Footer;
