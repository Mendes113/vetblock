// src/App.tsx
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';

import CreateConsultation from './pages/CreateConsult'; // Importando o componente de criação de consulta
import { Home } from './pages/Home'; // Página inicial
import { Navbar } from './components/Navbar/navbar';
import { Footer } from './components/Footer/Footer';
import PatientsPage from './pages/Patients';
import UserPage from './pages/UserPage';
import Login from './pages/Login';
import PrivateRoute from './components/PrivateRoute/PrivateRoute';

function App() {
  return (
    <div className="min-h-screen flex flex-col ">
      <Router>
        {/* Navbar no topo */}
        <header>
          <Navbar />
        </header>

        {/* Conteúdo principal da página */}
        <div className="flex-grow  w-[1500px] align-middle mx-auto">
          <Routes>
            {/* Rota da página de login */}
            <Route path="/login" element={<Login />} />
            
            {/* Rotas públicas */}
            <Route path="/" element={<Home />} />

            {/* Rotas protegidas */}
            <Route element={<PrivateRoute />}>
              <Route path="/create-consultation" element={<CreateConsultation onSave={() => {}} />} />
              <Route path="/patients-search" element={<PatientsPage />} />
              <Route path="/users" element={<UserPage />} />
            </Route>
            
            {/* Outras rotas podem ser adicionadas aqui */}
          </Routes>
        </div>
        {/* Footer no final */}
      
      </Router>
      <Footer />
    </div>
  );
}

export default App;
