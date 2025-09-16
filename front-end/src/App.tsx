import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navbar from '@/components/layout/Navbar/Navbar';
import Footer from '@/components/layout/Footer/Footer';
import Clients from '@/pages/Clients/Clients';
import './App.css';
import './index.css';

function App() {
  return (
    <Router>
      <div className="App" style={{ display: 'flex', flexDirection: 'column', minHeight: '100vh' }}>
        {/* Main Application Structure */}
        <Navbar />

        {/* Main Content */}
        <main className="container" style={{ flex: 1 }}>
          <Routes>
            <Route
              path="/"
              element={
                <div className="hero">
                  <h1>Figueiredo Imobiliária</h1>
                  <p>Sistema de gestão imobiliária moderno e eficiente</p>
                  <button
                    onClick={() => (window.location.href = '/properties')}
                    className="btn btn-primary"
                  >
                    Ver Propriedades
                  </button>
                </div>
              }
            />
            <Route
              path="/properties"
              element={
                <div className="container">
                  <h1>Propriedades</h1>
                  <p>Lista de propriedades disponíveis</p>
                </div>
              }
            />
            <Route path="/clients" element={<Clients />} />
            <Route
              path="/leads"
              element={
                <div className="container">
                  <h1>Leads</h1>
                  <p>Gerenciamento de leads</p>
                </div>
              }
            />
            <Route
              path="/stats/darwin"
              element={
                <div className="container">
                  <h1>Estatísticas Darwin</h1>
                  <p>Análise de conversões</p>
                </div>
              }
            />
            <Route
              path="/stats/company"
              element={
                <div className="container">
                  <h1>Estatísticas da Empresa</h1>
                  <p>Desempenho da equipe</p>
                </div>
              }
            />
            <Route
              path="/finances"
              element={
                <div className="container">
                  <h1>Finanças</h1>
                  <p>Controle financeiro</p>
                </div>
              }
            />
            <Route
              path="/profile"
              element={
                <div className="container">
                  <h1>Perfil</h1>
                  <p>Configurações do usuário</p>
                </div>
              }
            />
            <Route
              path="/settings"
              element={
                <div className="container">
                  <h1>Configurações</h1>
                  <p>Configurações do sistema</p>
                </div>
              }
            />
            <Route
              path="/search"
              element={
                <div className="container">
                  <h1>Busca</h1>
                  <p>Resultados da busca</p>
                </div>
              }
            />
          </Routes>
        </main>
        
        {/* Footer */}
        <Footer />
      </div>
    </Router>
  );
}

export default App;
