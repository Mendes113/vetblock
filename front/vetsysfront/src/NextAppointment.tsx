import React, { useState, useEffect } from 'react';
import { Button } from './components/ui/button';
import { Skeleton } from './components/ui/skeleton';
import PatientModal from './components/patientModal/patientmodal';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCalendarCheck } from '@fortawesome/free-solid-svg-icons';
import { Patient, Consultation } from './Types';
import { FiClock, FiEdit3, FiClipboard } from 'react-icons/fi';
// Defina o tipo Patient com todas as propriedades necessárias para ambos os contextos



const NextAppointment: React.FC<{ appointment: Consultation | null }> = ({ appointment }) => {
  if (!appointment) {
    return <div className="p-4 bg-gray-200 rounded-lg">Nenhum compromisso agendado.</div>;
  }

  const [isModalOpen, setIsModalOpen] = useState(false);
  const [pacienteData, setPacienteData] = useState<Patient | null>(null); // Tipagem correta para Patient
  const [isLoading, setIsLoading] = useState(true); // Estado de carregamento

  useEffect(() => {
    const fetchAnimal = async (id: string) => {
      try {
        const response = await fetch(`http://localhost:8081/api/v1/animals/${id}`);
        if (!response.ok) {
          throw new Error("Erro ao buscar o animal");
        }
        const dataAnimal = await response.json();
        
        // Após buscar o animal, busque as consultas
        const consultationsResponse = await fetch(`http://localhost:8081/api/v1/consultations/patient/${id}`);
        if (!consultationsResponse.ok) {
          throw new Error("Erro ao buscar as consultas");
        }
        const dataConsultations = await consultationsResponse.json();
        const imagePrefix = dataAnimal.image.startsWith('data:image/')
          ? ''
          : 'data:image/png;base64,';
        //add prefixo para imagem
        dataAnimal.image = imagePrefix + dataAnimal.image;

    
        console.log(dataAnimal.image.includes('data:image/png;base64,'));
        setPacienteData({
          ...dataAnimal,
          photoUrl: `${dataAnimal.image}`,
          lastConsultations: dataConsultations, // Atualiza com as consultas reais
          lastMedications: [], // Definir isso com dados reais quando necessário
        });
       
        setIsLoading(false); // Finaliza o carregamento
      } catch (error) {
        console.error("Erro ao buscar dados:", error);
        setIsLoading(false); // Finaliza o carregamento em caso de erro
      }
      
    };

    if (appointment) {
      fetchAnimal(appointment.animal_id); // Usa o `appointment.animal_id` corretamente
    }

  }, [appointment]);


  
  const openModal = () => {
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
  };

  const { consultation_date: date, consultation_hour, observation, reason } = appointment;
  return (
    <div className="p-6 bg-white shadow-lg rounded-lg border border-gray-200 max-w-3xl mx-auto">
      <h2 className="text-xl font-semibold mb-4 text-gray-800 flex items-center gap-2">
        Próximo Compromisso <FontAwesomeIcon icon={faCalendarCheck} className="text-green-500" />
      </h2>

      {/* Informações do compromisso */}
      <div className="flex flex-col sm:flex-row flex-wrap items-center justify-between gap-6 mt-4">
        <p className="text-gray-700 flex items-center gap-2">
          <FiClipboard className="text-gray-500" />
          <strong>Data:</strong> {new Date(date).toLocaleDateString()}
        </p>
        <p className="text-gray-700 flex items-center gap-2">
          <FiClock className="text-gray-500" />
          <strong>Hora:</strong> {consultation_hour}
        </p>
        <p className="text-gray-700 flex items-center gap-2">
          <FiEdit3 className="text-gray-500" />
          <strong>Motivo:</strong> {reason}
        </p>
        <p className="text-gray-700 flex items-center gap-2">
          <FiEdit3 className="text-gray-500" />
          <strong>Observação:</strong> {observation}
        </p>
      </div>

     {/* Paciente */}
<div className="mt-6">
  <h3 className="text-lg font-semibold mb-4 text-gray-800">Paciente</h3>

  {/* Card para Paciente */}
  <div className="flex flex-col sm:flex-row items-center p-4 bg-gray-50 rounded-lg shadow-sm gap-6">
    {isLoading ? (
      <Skeleton className="w-24 h-24 rounded-full" />
    ) : pacienteData?.photoUrl ? (
      <img
        src={pacienteData.photoUrl.startsWith('data:image/')
          ? pacienteData.photoUrl
          : `data:image/png;base64,${pacienteData.photoUrl}`}
        alt={`Foto de ${pacienteData.name}`}
        className="w-24 h-24 rounded-full object-cover"
      />
    ) : (
      <Skeleton className="w-24 h-24 rounded-full" />
    )}

    <div className="grid grid-cols-2 sm:grid-cols-3 gap-4 w-full">
      <div className="flex flex-col">
        <p className="text-gray-600 font-bold">Nome do Paciente</p>
        {isLoading ? (
          <Skeleton className="w-20 h-4" />
        ) : (
          <p className="text-gray-700">{pacienteData?.name}</p>
        )}
      </div>

      <div className="flex flex-col">
        <p className="text-gray-600 font-bold">Raça</p>
        {isLoading ? (
          <Skeleton className="w-20 h-4" />
        ) : (
          <p className="text-gray-700">{pacienteData?.breed}</p>
        )}
      </div>

      <div className="flex flex-col">
        <p className="text-gray-600 font-bold">Idade</p>
        {isLoading ? (
          <Skeleton className="w-20 h-4" />
        ) : (
          <p className="text-gray-700">{pacienteData?.age} anos</p>
        )}
      </div>

      <div className="flex flex-col">
        <p className="text-gray-600 font-bold">Peso</p>
        {isLoading ? (
          <Skeleton className="w-20 h-4" />
        ) : (
          <p className="text-gray-700">{pacienteData?.weight || 'N/A'} kg</p>
        )}
      </div>

      <div className="flex flex-col">
        <p className="text-gray-600 font-bold">Espécie</p>
        {isLoading ? (
          <Skeleton className="w-20 h-4" />
        ) : (
          <p className="text-gray-700">{pacienteData?.species}</p>
        )}
      </div>
    </div>

    <div className="mt-4 sm:mt-0">
      <Button className="bg-blue-500 text-white px-4 py-2 rounded-md hover:bg-blue-600" onClick={openModal}>
        Ver Paciente
      </Button>
    </div>

    {isModalOpen && (
      <PatientModal patient={pacienteData} loading={isLoading} onClose={closeModal} />
    )}
  </div>
</div>

    </div>
  );
};

export default NextAppointment;