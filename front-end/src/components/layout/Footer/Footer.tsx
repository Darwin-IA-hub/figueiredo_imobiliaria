import React from 'react';
import './Footer.css';

const Footer: React.FC = () => {
  return (
    <footer className="footer">
      <div className="footer-container">
        <div className="footer-brand">
          <h2>Figueiredo Imobiliária</h2>
          <p>Desde março de 1996 — quase 30 anos</p>
          <p>Imóvel, moeda forte. Aceitamos sugestões de novo slogan.</p>
        </div>

        <div className="footer-content">
          <div className="footer-section">
            <h3>Consultores</h3>
            <ul>
              <li>Segunda a Sexta-feira</li>
              <li>Horário: 9h às 17h</li>
              <li>Intervalo: 12h às 14h</li>
            </ul>
          </div>

          <div className="footer-section">
            <h3>Contato</h3>
            <ul>
              <li>Site: <a href="https://figueiredoimoveis.com.br" target="_blank" rel="noreferrer">figueiredoimoveis.com.br</a></li>
              <li>E-mail: <a href="mailto:Contato@figueiredoimoveis.com.br">Contato@figueiredoimoveis.com.br</a> <span>(confirmar com @Robison DJ?)</span></li>
              <li>Telefone: <a href="tel:+5515981517070">15 98151-7070</a></li>
            </ul>
          </div>

          <div className="footer-section">
            <h3>Sede</h3>
            <p>Av. Cel. Firmo Vieira de Camargo, 825, Centro — Tatuí/SP</p>
          </div>
        </div>
      </div>
    </footer>
  );
};

export default Footer;
