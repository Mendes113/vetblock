import React, { useState, useEffect } from 'react';
import { Card, CardContent } from '../components/ui/card';
import { Button } from '../components/ui/button';
import axios from 'axios'; // Assumindo que está usando axios para requisições

interface CreateConsultationProps {
  onSave: (consultation: NewConsultation) => void;
}

interface NewConsultation {
  animalName: string;
  species: string;
  breed: string;
  age: string;
  weight: string;
  consultationDate: string;
  consultationTime: string;
  reason: string;
  notifyOwner?: boolean;
}

interface Animal {
  id: string;
  name: string;
  species: string;
  breed: string;
  age: string;
  weight: string;
}

const CreateConsultation: React.FC<CreateConsultationProps> = ({ onSave }) => {
  const [animalName, setAnimalName] = useState('');
  const [species, setSpecies] = useState('');
  const [breed, setBreed] = useState('');
  const [age, setAge] = useState('');
  const [weight, setWeight] = useState('');
  const [consultationDate, setConsultationDate] = useState('');
  const [consultationTime, setConsultationTime] = useState('');
  const [reason, setReason] = useState('');
  const [notifyOwner, setNotifyOwner] = useState(false);
  const [animalSuggestions, setAnimalSuggestions] = useState<Animal[]>([]);
  const [isLoading, setIsLoading] = useState(false);

  // Função para buscar animais cadastrados
  const fetchAnimals = async (query: string) => {
    if (query.length < 2) {
      setAnimalSuggestions([]);
      return;
    }
    setIsLoading(true);
    try {
      const response = await axios.get(`/api/animals?name=${query}`);
      setAnimalSuggestions(response.data as Animal[]);
    } catch (error) {
      console.error("Erro ao buscar animais:", error);
    } finally {
      setIsLoading(false);
    }
  };

  // Atualiza o estado e busca animais ao digitar no campo de nome
  const handleAnimalNameChange = (value: string) => {
    setAnimalName(value);
    fetchAnimals(value);
  };

  // Seleciona um animal sugerido e preenche os campos
  const handleSelectAnimal = (animal: Animal) => {
    setAnimalName(animal.name);
    setSpecies(animal.species);
    setBreed(animal.breed);
    setAge(animal.age);
    setWeight(animal.weight);
    setAnimalSuggestions([]); // Limpa sugestões após seleção
  };

  const handleSave = () => {
    if (!animalName || !consultationDate || !reason) {
      alert('Preencha todos os campos obrigatórios!');
      return;
    }

    const newConsultation: NewConsultation = {
      animalName,
      species,
      breed,
      age,
      weight,
      consultationDate,
      consultationTime,
      reason,
      notifyOwner,
    };

    onSave(newConsultation);
  };

  return (
    <div className="container mx-auto p-6">
      <header className="bg-blue-600 text-white p-4 rounded-lg mb-6">
        <h1 className="text-3xl font-bold text-center">Agendar Nova Consulta</h1>
      </header>

      <Card className="w-full max-w-3xl mx-auto p-4 bg-white rounded-lg shadow-lg">
        <CardContent className="p-6">
          <h2 className="text-2xl font-bold mb-4 text-center">Agendar Nova Consulta</h2>

          {/* Nome do Animal com Autocomplete */}
          <div className="mb-4 relative">
            <label className="block text-gray-700 font-semibold mb-2" htmlFor="animalName">
              Nome do Animal
            </label>
            <input
              id="animalName"
              type="text"
              className="w-full px-3 py-2 border rounded-lg"
              value={animalName}
              onChange={(e) => handleAnimalNameChange(e.target.value)}
              placeholder="Digite o nome do animal"
            />
            {/* Sugestões de Animais */}
            {animalSuggestions.length > 0 && (
              <ul className="absolute bg-white border w-full rounded-lg mt-1 shadow-lg z-10">
                {animalSuggestions.map((animal) => (
                  <li
                    key={animal.id}
                    className="p-2 hover:bg-gray-200 cursor-pointer"
                    onClick={() => handleSelectAnimal(animal)}
                  >
                    {animal.name} - {animal.species}
                  </li>
                ))}
              </ul>
            )}
            {isLoading && <div className="text-gray-500">Carregando...</div>}
          </div>

          {/* Outros campos para a consulta */}
          <div className="mb-4">
            <label className="block text-gray-700 font-semibold mb-2" htmlFor="species">
              Espécie
            </label>
            <input
              id="species"
              type="text"
              className="w-full px-3 py-2 border rounded-lg"
              value={species}
              onChange={(e) => setSpecies(e.target.value)}
              placeholder="Ex: Cachorro, Gato, Coelho"
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700 font-semibold mb-2" htmlFor="breed">
              Raça
            </label>
            <input
              id="breed"
              type="text"
              className="w-full px-3 py-2 border rounded-lg"
              value={breed}
              onChange={(e) => setBreed(e.target.value)}
              placeholder="Digite a raça do animal"
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700 font-semibold mb-2" htmlFor="age">
              Idade
            </label>
            <input
              id="age"
              type="text"
              className="w-full px-3 py-2 border rounded-lg"
              value={age}
              onChange={(e) => setAge(e.target.value)}
              placeholder="Digite a idade do animal"
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700 font-semibold mb-2" htmlFor="weight">
              Peso (kg)
            </label>
            <input
              id="weight"
              type="text"
              className="w-full px-3 py-2 border rounded-lg"
              value={weight}
              onChange={(e) => setWeight(e.target.value)}
              placeholder="Peso do animal em kg"
            />
          </div>

          {/* Opção para notificar */}
          <div className="mb-4">
            <input
              id="notifyOwner"
              type="checkbox"
              checked={notifyOwner}
              onChange={(e) => setNotifyOwner(e.target.checked)}
            />
            <label htmlFor="notifyOwner" className="text-gray-700 ml-2">
              Enviar lembrete para o responsável por SMS/Email
            </label>
          </div>

        
          <div className="flex justify-end gap-4">
            <Button className="bg-gray-400 text-white px-4 py-2 rounded-lg hover:bg-gray-500" onClick={() => window.history.back()}>
              Cancelar
            </Button>
            <Button className="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700" onClick={handleSave}>
              Salvar Consulta
            </Button>
          </div>
        </CardContent>
      </Card>
    </div>
  );
};

export default CreateConsultation;