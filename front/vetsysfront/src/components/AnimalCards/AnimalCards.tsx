import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Card, CardContent } from '../ui/card';
import { Button } from '../ui/button';
import PatientModal from '../patientModal/patientmodal'; // Certifique-se de ajustar o caminho conforme necessário

interface Animal {
  animal_id: string;
  name: string;
  species: string;
  breed: string;
  age: number;
  weight: number;
  image: string;
  description: string;
}

interface AnimalCardsProps {
  searchTerm: string;
  currentPage: number;
  onPageChange: (newPage: number) => void;
}

const AnimalCards: React.FC<AnimalCardsProps> = ({ searchTerm, currentPage, onPageChange }) => {
  const [animals, setAnimals] = useState<Animal[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [totalPages, setTotalPages] = useState(1);
  const [selectedAnimal, setSelectedAnimal] = useState<Animal | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const itemsPerPage = 12;

  useEffect(() => {
    fetchAnimals();
  }, [searchTerm, currentPage]);

  const fetchAnimals = async () => {
    setIsLoading(true);
    try {
      const response = await axios.get<Animal[]>('http://localhost:8081/api/v1/animals');
      const filteredAnimals = response.data.filter((animal) =>
        animal.name.toLowerCase().includes(searchTerm.toLowerCase())
      );
      
      const startIndex = (currentPage - 1) * itemsPerPage;
      const paginatedAnimals = filteredAnimals.slice(startIndex, startIndex + itemsPerPage);

      setAnimals(paginatedAnimals);
      setTotalPages(Math.ceil(filteredAnimals.length / itemsPerPage));
    } catch (error) {
      console.error('Erro ao buscar animais:', error);
    } finally {
      setIsLoading(false);
    }
  };

  const handleOpenModal = (animal: Animal) => {
    setSelectedAnimal(animal);
    setIsModalOpen(true);
  };

  const handleCloseModal = () => {
    setSelectedAnimal(null);
    setIsModalOpen(false);
  };

  const handlePrevPage = () => {
    if (currentPage > 1) {
      onPageChange(currentPage - 1);
    }
  };

  const handleNextPage = () => {
    if (currentPage < totalPages) {
      onPageChange(currentPage + 1);
    }
  };

  return (
    <div className="p-4">
      {isLoading ? (
        <div className="flex justify-center items-center text-gray-700">
          <div className="animate-spin rounded-full h-8 w-8 border-t-2 border-b-2 border-gray-500 mr-2"></div>
          Carregando pacientes...
        </div>
      ) : animals.length === 0 ? (
        <div className="text-center text-gray-700">Nenhum paciente encontrado.</div>
      ) : (
        <>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {animals.map((animal) => (
              <Card 
                key={animal.animal_id} 
                className="bg-white rounded-xl shadow-md hover:shadow-lg transition-shadow duration-300"
              >
                <CardContent>
                  <img
                    src={animal.image !== "00000000-0000-0000-0000-000000000000" ? `/path/to/images/${animal.image}` : "https://cdn.discordapp.com/attachments/1058806818224222278/1296986375995916378/9k.png?ex=67144830&is=6712f6b0&hm=b824f7eaba712aaf03e7fd0fe7d5890eab143e84d90de03ad632a6cb837b38a5&"}
                    alt={animal.name}
                    className="w-full h-48 object-cover rounded-t-lg"
                  />
                  <div className="p-4">
                    <h2 className="text-xl font-semibold mb-2 text-gray-800">{animal.name}</h2>
                    <p className="text-gray-600 mb-1"><strong>Espécie:</strong> {animal.species}</p>
                    <p className="text-gray-600 mb-1"><strong>Raça:</strong> {animal.breed}</p>
                    <p className="text-gray-600 mb-1"><strong>Idade:</strong> {animal.age} anos</p>
                    <p className="text-gray-600 mb-1"><strong>Peso:</strong> {animal.weight} kg</p>
                    <p className="text-gray-600 mb-3"><strong>Descrição:</strong> {animal.description}</p>
                    <Button 
                      className="bg-green-500 text-white px-4 py-2 rounded-lg hover:bg-green-600 transition-colors duration-200"
                      onClick={() => handleOpenModal(animal)}
                    >
                      Ver Detalhes
                    </Button>
                  </div>
                </CardContent>
              </Card>
            ))}
          </div>
          
          {/* Controles de Paginação */}
          <div className="flex justify-center mt-8">
            <Button 
              className={`px-4 py-2 rounded-lg mx-2 ${currentPage === 1 ? 'bg-gray-300 cursor-not-allowed' : 'bg-gray-400 text-white hover:bg-gray-500'} transition-all duration-200`} 
              onClick={handlePrevPage}
              disabled={currentPage === 1}
            >
              Anterior
            </Button>
            <Button 
              className={`px-4 py-2 rounded-lg mx-2 ${currentPage === totalPages ? 'bg-gray-300 cursor-not-allowed' : 'bg-gray-400 text-white hover:bg-gray-500'} transition-all duration-200`} 
              onClick={handleNextPage}
              disabled={currentPage === totalPages}
            >
              Próximo
            </Button>
          </div>
          
          {/* Modal Paciente */}
          {isModalOpen && selectedAnimal && (
            <PatientModal
              patient={selectedAnimal}
              loading={false} // Ajuste conforme necessário
              onClose={handleCloseModal}
            />
          )}
        </>
      )}
    </div>
  );
};

export default AnimalCards;
