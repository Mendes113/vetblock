import React, { useState } from 'react';
import AnimalCards from '../components/AnimalCards/AnimalCards'; // Ajuste o caminho conforme necessário

const PatientsPage: React.FC = () => {
  const [searchTerm, setSearchTerm] = useState('');
  const [currentPage, setCurrentPage] = useState(1);

  const handleSearchChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(e.target.value);
    setCurrentPage(1); // Resetar para a primeira página ao buscar
  };

  const handlePageChange = (newPage: number) => {
    setCurrentPage(newPage);
  };

  return (
    <div className="min-h-screen bg-gray-100">
      <header className="bg-blue-600 text-white p-4 rounded-lg mb-6">
        <h1 className="text-3xl font-bold text-center">Pacientes Cadastrados</h1>
      </header>
      
      <div className="container mx-auto p-6">
        <div className="mb-4">
          <input
            type="text"
            className="w-full px-3 py-2 border rounded-lg"
            placeholder="Buscar por nome do paciente..."
            value={searchTerm}
            onChange={handleSearchChange}
          />
        </div>

        <AnimalCards 
          searchTerm={searchTerm} 
          currentPage={currentPage} 
          onPageChange={handlePageChange} 
        />
      </div>
    </div>
  );
};

export default PatientsPage;
