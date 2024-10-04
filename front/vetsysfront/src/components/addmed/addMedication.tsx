import React, { useState } from 'react';
import { Button } from '../ui/button';
import { Card, CardContent } from '../ui/card';
import { Input } from '../ui/input'; // Assumindo que existe um componente de input
import { v4 as uuidv4 } from 'uuid';

interface Patient {
  id: string;
  name: string;
  species: string;
  breed: string;
  age: string;
  weight: string;
  description: string;
  lastConsultations: { id: string; date: string; reason: string }[];
  lastMedications: string[];
}

interface NewMedication {
  id: string;
  name: string;
  description: string;
  price: number;
  batchNumber: string;
  concentration: string;
  presentation: string;
  dosageForm: string;
  activePrinciples: string[];
  manufacturer: string;
  quantity: number;
  unit: string;
  storageConditions: string;
  prescriptionRequired: boolean;
  expiration: Date;
}

interface NewDosage {
  id: string;
  medicationId: string;
  animalId: string;
  startDate: Date;
  endDate: Date;
  quantity: number;
  dosage: string;
  administrationRoute: string;
  frequency: string;
  durationDays: number;
}




interface AddMedicationCardProps {
  onClose: () => void;
  onAddMedication: (medication: NewMedication, dosage: NewDosage) => void;

  patient: Patient; // Informações do paciente
}

export const AddMedicationCard: React.FC<AddMedicationCardProps> = ({ onClose, onAddMedication, patient }) => {
  const [activeTab, setActiveTab] = useState<'info' | 'medication' | 'history'>('info'); // Estado para gerenciar a aba ativa

  const [medication, setMedication] = useState<NewMedication>({
    id: uuidv4(),
    name: '',
    description: '',
    price: 0,
    batchNumber: '',
    concentration: '',
    presentation: '',
    dosageForm: '',
    activePrinciples: [],
    manufacturer: '',
    quantity: 0,
    unit: '',
    storageConditions: '',
    prescriptionRequired: false,
    expiration: new Date(),
  });

  
  const [dosage, setDosage] = useState<NewDosage>({
    id: uuidv4(),
    medicationId: medication.id,
    animalId: uuidv4(), // Substituir pelo ID correto do animal
    startDate: new Date(),
    endDate: new Date(),
    quantity: 1,
    dosage: '',
    administrationRoute: '',
    frequency: '',
    durationDays: 0,
  });

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setMedication({ ...medication, [name]: value });
  };

  const handleDosageChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
    const { name, value } = e.target;
    setDosage({ ...dosage, [name]: value });
  };

  // Função para calcular a dosagem sugerida baseada no peso do animal
  const calculateSuggestedDose = (dosePerKg: number) => {
    return dosePerKg * parseFloat(patient.weight)
  };

  const handleAddMedication = () => {
    if (medication.name && dosage.dosage) {
      onAddMedication(medication, dosage);
      onClose();
    } else {
      alert('Preencha todos os campos obrigatórios.');
    }
  };

  return (
    <Card className="fixed inset-0 z-50 flex items-center justify-center bg-gray-900 bg-opacity-50 p-6">
      <CardContent className="w-full max-w-2xl bg-white p-6 rounded-lg shadow-lg">
        <h2 className="text-2xl font-bold mb-6 text-gray-800">Gerenciamento de Medicação</h2>

        {/* Tabs */}
        <div className="mb-6 border-b border-gray-200">
          <ul className="flex flex-wrap -mb-px text-sm font-medium text-center text-gray-500">
            <li className="mr-2">
                        <Button
            className={`inline-block p-4 rounded-t-lg border-b-2 transition-all duration-300 ease-in-out ${
              activeTab === 'medication'
              ? 'bg-gray-300 border-gray-500 translate-y-1 shadow-inner' // Botão pressionado
              : 'bg-gray-500 hover:bg-gray-200 shadow-lg hover:shadow-md'
            } shadow-lg hover:shadow-md`}
              onClick={() => setActiveTab('info')}
            >
              Informações do Paciente
            </Button>

            <Button
              className={`inline-block p-4 rounded-t-lg border-b-2 transition-all duration-300 ease-in-out ${
                activeTab === 'info'
                  ? 'border-gray-500 translate-y-1 shadow-inner' // Botão pressionado
                  : 'bg-gray-500 hover:bg-gray-200 shadow-lg hover:shadow-md'
              } shadow-lg hover:shadow-md`}
              onClick={() => setActiveTab('medication')}
            >
              Adicionar Medicação
            </Button>

            <Button
              className={`inline-block p-4 rounded-t-lg border-b-2 transition-all duration-300 ease-in-out ${
                activeTab === 'history'
                  ?  ' bg-gray-500 border-gray-500 translate-y-1 shadow-inner' // Botão pressionado
                  : ' hover:bg-gray-200 shadow-lg hover:shadow-md'
              }`}
              onClick={() => setActiveTab('history')}
            >
              Historico Paciente
            </Button>


            </li>
          </ul>
        </div>

        {/* Conteúdo da Tab */}
        {activeTab === 'info' && (
          <div>
            <h3 className="text-xl font-semibold mb-4">Informações do Paciente</h3>
            <p><strong>Nome:</strong> {patient.name}</p>
            <p><strong>Espécie:</strong> {patient.species}</p>
            <p><strong>Raça:</strong> {patient.breed}</p>
            <p><strong>Idade:</strong> {patient.age}</p>
            <p><strong>Peso:</strong> {patient.weight} kg</p>
            <h4 className="text-lg font-semibold mt-4">Últimas Consultas</h4>
            <ul>
              {patient.lastConsultations.map((consultation) => (
                <li key={consultation.id}>
                  <p>{consultation.date}: {consultation.reason}</p>
                </li>
              ))}
            </ul>
          </div>
        )}

        {activeTab === 'medication' && (
          <div>
            {/* Formulário de medicação */}
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
              <div className="col-span-2">
                <Input
                  name="name"
                  value={medication.name}
                  onChange={handleInputChange}
                  required
                  placeholder="Nome do Medicamento"
                  className="w-full"
                />
              </div>
              <div className="col-span-2">
                <Input
                  name="description"
                  value={medication.description}
                  onChange={handleInputChange}
                  placeholder="Descrição"
                  className="w-full"
                />
              </div>
              <Input
                name="price"
                type="number"
                value={medication.price}
                onChange={handleInputChange}
                required
                placeholder="Preço (R$)"
              />
              <Input
                name="concentration"
                value={medication.concentration}
                onChange={handleInputChange}
                required
                placeholder="Concentração (mg/ml, etc.)"
              />
            </div>

            {/* Formulário de dosagem */}
            <h3 className="text-lg font-semibold mt-4 mb-2">Dosagem</h3>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              <Input
                name="dosage"
                value={dosage.dosage}
                onChange={handleDosageChange}
                required
                placeholder="Dosagem (em mg)"
              />
              <select
                name="administrationRoute"
                value={dosage.administrationRoute}
                onChange={handleDosageChange}
                required
                className="w-full p-2 border border-gray-300 rounded-lg"
              >
                <option value="">Selecione a Via de Administração</option>
                <option value="oral">Oral</option>
                <option value="IV">Intravenosa</option>
                <option value="IM">Intramuscular</option>
                <option value="SC">Subcutânea</option>
              </select>

              <Input
                name="frequency"
                value={dosage.frequency}
                onChange={handleDosageChange}
                required
                placeholder="Frequência (ex: 1x ao dia)"
              />
              <Input
                name="durationDays"
                type="number"
                value={dosage.durationDays}
                onChange={handleDosageChange}
                required
                placeholder="Duração do tratamento (dias)"
              />
            </div>

            {/* Cálculo da dosagem sugerida */}
            <div className="mt-4">
              <label className="block text-sm font-medium text-gray-700">
                Dose recomendada por kg (mg/kg)
              </label>
              <Input
                type="number"
                onChange={(e) => setDosage({ ...dosage, dosage: calculateSuggestedDose(parseFloat(e.target.value)).toString() })}
                placeholder="Insira a dose recomendada por kg"
                className="w-full"
              />
              <p className="mt-2 text-gray-600">Dose sugerida: {dosage.dosage} mg</p>
            </div>
          </div>
        )}


{activeTab === 'history' && (
  <div className="p-4 bg-white rounded-lg shadow-md">
    <h2 className="text-2xl font-semibold mb-4">Histórico do Paciente</h2>
    
    {/* Dados do paciente */}
    <div className="mb-6">
      <h3 className="text-lg font-medium">Informações do Paciente</h3>
      <p><strong>Nome:</strong> Rex</p>
    </div>

    {/* Histórico de Consultas */}
    <div className="overflow-x-auto">
      <h3 className="text-lg font-medium mb-3">Consultas</h3>
      <table className="min-w-full bg-white border border-gray-200">
        <thead>
          <tr>
            <th className="px-4 py-2 border-b">Data</th>
            <th className="px-4 py-2 border-b">Diagnóstico</th>
            <th className="px-4 py-2 border-b">Tratamento</th>
          </tr>
        </thead>
        <tbody>
          {/* Exemplo de dados de consultas */}
          <tr>
            <td className="px-4 py-2 border-b">2023-09-01</td>
            <td className="px-4 py-2 border-b">Consulta de rotina</td>
            <td className="px-4 py-2 border-b">Vermifugação</td>
          </tr>
          <tr>
            <td className="px-4 py-2 border-b">2023-08-01</td>
            <td className="px-4 py-2 border-b">Infecção de ouvido</td>
            <td className="px-4 py-2 border-b">Antibiótico</td>
          </tr>
          <tr>
            <td className="px-4 py-2 border-b">2023-07-15</td>
            <td className="px-4 py-2 border-b">{patient.lastMedications[0]}</td>
            <td className="px-4 py-2 border-b">Vacina antirrábica</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
)}


        {/* Botões de ação */}
        <div className="flex justify-end gap-4 mt-6">
          <Button
            className="bg-gray-600 text-white px-4 py-2 rounded-lg hover:bg-gray-700 transition duration-150 ease-in-out"
            onClick={onClose}
          >
            Cancelar
          </Button>
          <Button
            className="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition duration-150 ease-in-out"
            onClick={handleAddMedication}
          >
            Salvar Medicação
          </Button>
        </div>
      </CardContent>
    </Card>
  );
};
