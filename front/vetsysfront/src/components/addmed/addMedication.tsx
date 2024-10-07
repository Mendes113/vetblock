import React, { useState } from 'react';
import { Button } from '../ui/button';
import { Card, CardContent } from '../ui/card';
import { Input } from '../ui/input'; // Assumindo que existe um componente de input
import { v4 as uuidv4 } from 'uuid';
import { FiUser, FiCalendar, FiList,  FiDollarSign, FiPackage, FiZap, FiClipboard, FiCheckCircle   } from 'react-icons/fi';
import { FaPrescriptionBottleAlt } from 'react-icons/fa';
import { IoTimeSharp } from "react-icons/io5";
import { MdOutlineDescription } from 'react-icons/md';
import { HiMiniCalendarDays } from "react-icons/hi2";
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
  const [activeTab, setActiveTab] = useState<'info' | 'medication' | 'history'>('info');
  
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
    animalId: uuidv4(),
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

  const calculateSuggestedDose = (dosePerKg: number) => {
    return dosePerKg * parseFloat(patient.weight);
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
                className={`inline-block p-4 rounded-t-lg transition-all duration-300 ease-in-out ${
                  activeTab === 'info'
                    ? 'bg-gray-300 border-gray-500 shadow-inner translate-y-1'
                    : 'bg-gray-500 hover:bg-gray-200 shadow-lg hover:shadow-md'
                }`}
                onClick={() => setActiveTab('info')}
              >
                Informações do Paciente
              </Button>

              <Button
                className={`inline-block p-4 rounded-t-lg transition-all duration-300 ease-in-out ${
                  activeTab === 'medication'
                    ? 'bg-gray-300 border-gray-500 shadow-inner translate-y-1'
                    : 'bg-gray-500 hover:bg-gray-200 shadow-lg hover:shadow-md'
                }`}
                onClick={() => setActiveTab('medication')}
              >
                Adicionar Medicação
              </Button>

              <Button
                className={`inline-block p-4 rounded-t-lg transition-all duration-300 ease-in-out ${
                  activeTab === 'history'
                    ? 'bg-gray-300 border-gray-500 shadow-inner translate-y-1'
                    : 'bg-gray-500 hover:bg-gray-200 shadow-lg hover:shadow-md'
                }`}
                onClick={() => setActiveTab('history')}
              >
                Histórico Paciente
              </Button>
            </li>
          </ul>
        </div>

        {/* Conteúdo da Tab */}
        {activeTab === 'info' && (
         <div className="p-6 bg-gray-50 rounded-lg shadow-lg">
         {/* Informações do Paciente */}
         <div className="mb-6">
           <h3 className="text-xl font-semibold text-gray-800 flex items-center gap-2 mb-4">
             <FiUser className="text-gray-600" />
             Informações do Paciente
           </h3>
           <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 text-gray-700">
             <p><strong>Nome:</strong> {patient.name}</p>
             <p><strong>Espécie:</strong> {patient.species}</p>
             <p><strong>Raça:</strong> {patient.breed}</p>
             <p><strong>Idade:</strong> {patient.age}</p>
             <p><strong>Peso:</strong> {patient.weight} kg</p>
           </div>
         </div>
       
         {/* Últimas Consultas */}
         <div className="mt-4">
           <h4 className="text-lg font-semibold text-gray-800 flex items-center gap-2">
             <FiCalendar className="text-gray-600" />
             Últimas Consultas
           </h4>
           <ul className="max-h-32 overflow-y-auto mt-2 grid grid-cols-1 gap-4">
             {patient.lastConsultations.map((consultation) => (
               <li key={consultation.id} className="bg-white p-4 rounded-lg shadow-md">
                 <p className="text-sm text-gray-600">
                   <strong>Data:</strong> {consultation.date}
                 </p>
                 <p className="text-sm text-gray-600">
                   <strong>Motivo:</strong> {consultation.reason}
                 </p>
               </li>
             ))}
           </ul>
         </div>
       </div>
        )}

        {activeTab === 'medication' && (
        


<div className="p-6 bg-gray-50 rounded-lg shadow-md">
  {/* Formulário de Medicação */}
  <h3 className="text-lg font-semibold text-gray-800 mb-4">Informações do Medicamento</h3>
  <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-6">
    <div className="relative col-span-2">
      <label className="block text-sm font-medium text-gray-600 absolute top-0 left-12 transform -translate-y-3 scale-90 transition-all duration-200 ease-in-out bg-gray-50 px-1">
        Nome do Medicamento
      </label>
      <Input
        name="name"
        value={medication.name}
        onChange={handleInputChange}
        required
        placeholder=" "
        className="w-full p-4 pl-12 rounded-lg border border-gray-300 shadow-sm focus:ring focus:ring-blue-200" // Ajuste no padding-left
      />
      <span className="absolute top-3 left-4 text-gray-400">
        <MdOutlineDescription />
      </span>
    </div>

    <div className="relative col-span-2">
      <label className="block text-sm font-medium text-gray-600 absolute top-0 left-12 transform -translate-y-3 scale-90 transition-all duration-200 ease-in-out bg-gray-50 px-1">
        Descrição
      </label>
      <Input
        name="description"
        value={medication.description}
        onChange={handleInputChange}
        placeholder=" "
        className="w-full p-4 pl-12 rounded-lg border border-gray-300 shadow-sm focus:ring focus:ring-blue-200"
      />
      <span className="absolute top-3 left-4 text-gray-400">
        <MdOutlineDescription />
      </span>
    </div>

    <div className="relative">
      <label className="block text-sm font-medium text-gray-600 absolute top-0 left-12 transform -translate-y-3 scale-90 transition-all duration-200 ease-in-out bg-gray-50 px-1">
        Preço (R$)
      </label>
      <Input
        name="price"
        type="number"
        value={medication.price}
        onChange={handleInputChange}
        required
        placeholder=" "
        className="w-full pl-12 p-4 rounded-lg border border-gray-300 shadow-sm focus:ring focus:ring-blue-200"
      />
      <span className="absolute top-3 left-4 text-gray-400">
        <FiDollarSign />
      </span>
    </div>

    <div className="relative">
      <label className="block text-sm font-medium text-gray-600 absolute top-0 left-12 transform -translate-y-3 scale-90 transition-all duration-200 ease-in-out bg-gray-50 px-1">
        Concentração (mg/ml, etc.)
      </label>
      <Input
        name="concentration"
        value={medication.concentration}
        onChange={handleInputChange}
        required
        placeholder=" "
        className="w-full pl-12 p-4 rounded-lg border border-gray-300 shadow-sm focus:ring focus:ring-blue-200"
      />
      <span className="absolute top-3 left-4 text-gray-400">
        <FiPackage />
      </span>
    </div>
  </div>

  {/* Formulário de Dosagem */}
  <h3 className="text-lg font-semibold text-gray-800 mt-6 mb-4">Dosagem</h3>
  <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
    <div className="relative">
      <label className="block text-sm font-medium text-gray-600 absolute top-0 left-12 transform -translate-y-3 scale-90 transition-all duration-200 ease-in-out bg-gray-50 px-1">
        Dosagem (em mg)
      </label>
      <Input
        name="dosage"
        value={dosage.dosage}
        onChange={handleDosageChange}
        required
        placeholder=" "
        className="w-full pl-12 p-4 rounded-lg border border-gray-300 shadow-sm focus:ring focus:ring-blue-200"
      />
      <span className="absolute top-3 left-4 text-gray-400">
        <FaPrescriptionBottleAlt />
      </span>
    </div>

    <div className="relative">
      <label className="block text-sm font-medium text-gray-600 absolute top-0 left-12 transform -translate-y-3 scale-90 transition-all duration-200 ease-in-out bg-gray-50 px-1">
        Frequência (ex: 1x ao dia)
      </label>
      <Input
        name="frequency"
        value={dosage.frequency}
        onChange={handleDosageChange}
        required
        placeholder=" "
        className="w-full p-4 pl-12 rounded-lg border border-gray-300 shadow-sm focus:ring focus:ring-blue-200"
      />
      <span className="absolute top-3 left-4 text-gray-400">
        <IoTimeSharp />
      </span>
    </div>

    <div className="relative">
      <label className="block text-sm font-medium text-gray-600 absolute top-0 left-12 transform -translate-y-3 scale-90 transition-all duration-200 ease-in-out bg-gray-50 px-1">
        Duração do tratamento (dias)
      </label>
      <Input
        name="durationDays"
        type="number"
        value={dosage.durationDays}
        onChange={handleDosageChange}
        required
        placeholder=" "
        className="w-full p-4 pl-12 rounded-lg border border-gray-300 shadow-sm focus:ring focus:ring-blue-200"
      />
      <span className="absolute top-3 left-4 text-gray-400">
        <HiMiniCalendarDays />
      </span>
    </div>
  </div>

  {/* Cálculo da dosagem sugerida */}
  <div className="mt-6">
    <label className="block text-sm font-medium text-gray-800 mb-2">
      Dose recomendada por kg (mg/kg)
    </label>
    <div className="relative">
      <Input
        type="number"
        onChange={(e) =>
          setDosage({
            ...dosage,
            dosage: calculateSuggestedDose(parseFloat(e.target.value)).toString(),
          })
        }
        placeholder=" "
        className="w-full pl-12 p-4 rounded-lg border border-gray-300 shadow-sm focus:ring focus:ring-blue-200"
      />
      <span className="absolute top-3 left-4 text-gray-400">
        <FiZap />
      </span>
    </div>
    <p className="mt-2 text-gray-600">Dose sugerida: {dosage.dosage} mg</p>
  </div>
</div>

        )}

        {activeTab === 'history' && (
        

         <div className="p-6 bg-white rounded-lg shadow-lg max-h-96 overflow-y-auto">
           <h2 className="text-2xl font-bold mb-6 text-gray-800">Histórico do Paciente</h2>
         
           {/* Informações do Paciente */}
           <div className="mb-6">
             <h3 className="text-lg font-medium text-gray-700">Informações do Paciente</h3>
             <p className="text-gray-600"><strong>Nome:</strong> {patient.name}</p>
           </div>
         
           {/* Histórico de Consultas */}
           <div>
             <h3 className="text-lg font-semibold text-gray-700 mb-4 flex items-center">
               <FiClipboard className="mr-2 text-gray-500" />
               Consultas
             </h3>
             <div className="space-y-4">
               {patient.lastConsultations.map((consultation, index) => (
                 <div key={index} className="flex items-center p-4 bg-gray-50 rounded-lg shadow-sm">
                   <div className="w-1/4 flex items-center">
                     <FiCalendar className="text-gray-500 mr-2" />
                     <span className="text-gray-600 text-sm">{consultation.date}</span>
                   </div>
                   <div className="w-2/4">
                     <p className="text-sm text-gray-700"><strong>Diagnóstico:</strong> {consultation.reason}</p>
                   </div>
                   <div className="w-1/4 flex items-center">
                     <FiCheckCircle className="text-green-500 mr-2" />
                     <span className="text-gray-600 text-sm">Tratamento</span>
                   </div>
                 </div>
               ))}
             </div>
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
