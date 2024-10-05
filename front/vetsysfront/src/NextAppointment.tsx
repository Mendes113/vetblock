import React, { useState, useEffect } from 'react';
import { Button } from './components/ui/button';
import { Skeleton } from './components/ui/skeleton';
import PatientModal from './components/patientModal/patientmodal';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCalendarCheck } from '@fortawesome/free-solid-svg-icons';
import { Patient, Consultation } from './Types';
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
    <div className="p-4 bg-white shadow-md rounded-lg border border-gray-300 w-[900px] justify-center items-center ml-auto mr-auto">
      <h2 className="text-lg font-semibold mb-2">
        Próximo Compromisso <FontAwesomeIcon icon={faCalendarCheck} />
      </h2>
      <div className="flex flex-col gap-4 items-center justify-center mt-5 sm:flex-row sm:gap-10">
        <p className="text-gray-700">
          <strong>Data:</strong> {new Date(date).toLocaleDateString()}
        </p>
        <p className="text-gray-700">
          <strong>Hora:</strong> {consultation_hour}
        </p>
        <p className="text-gray-700">
          <strong>Motivo:</strong> {reason}
        </p>
        <p className="text-gray-700">
          <strong>Observação:</strong> {observation}
        </p>
      </div>

      {/* paciente */}
      <div className="mt-4">
        <h3 className="text-lg font-semibold mb-2">Paciente</h3>
        {/* card for patient */}
        <div className="flex flex-col items-center p-4 bg-gray-100 rounded-lg [w-1200]:flex-row sm:gap-4">
          {isLoading ? (
            <Skeleton className="w-24 h-24 rounded-full" />
          ) : pacienteData?.photoUrl ? (
            <img
            src={pacienteData.photoUrl.startsWith('data:image/') 
              ? pacienteData.photoUrl 
              : `data:image/png;base64,${pacienteData.photoUrl}`} // Garantir o prefixo
            alt={`Foto de ${pacienteData.name}`}
            className="w-24 h-24 rounded-full object-cover"
          />

          
          ) : (
            <Skeleton className="w-24 h-24 rounded-full" />
          )}

          <div className='flex flex-col items-center text-center sm:flex-row sm:gap-8'>
            <div className='flex flex-col items-center'>
              <p className="text-gray-700 font-bold">Nome do Paciente</p>
              {isLoading ? (
                <Skeleton className="w-20 h-4 flex">
                  <div className="w-full h-full bg-gray-300 rounded"></div>
                </Skeleton>
              ) : (
                <p className="text-gray-700">{pacienteData?.name}</p>
              )}
            </div>

            <div className='flex flex-col items-center'>
              <p className="text-gray-700 font-bold">Raça</p>
              {isLoading ? (
                <Skeleton className="w-20 h-4 flex">
                  <div className="w-full h-full bg-gray-300 rounded"></div>
                </Skeleton>
              ) : (
                <p className="text-gray-700">{pacienteData?.breed}</p>
              )}
            </div>

            <div className='flex flex-col items-center'>
              <p className="text-gray-700 font-bold">Idade</p>
              {isLoading ? (
                <Skeleton className="w-20 h-4 flex">
                  <div className="w-full h-full bg-gray-300 rounded"></div>
                </Skeleton>
              ) : (
                <p className="text-gray-700">{pacienteData?.age} anos</p>
              )}
            </div>

            <div className='flex flex-col items-center'>
              <p className="text-gray-700 font-bold">Peso</p>
              {isLoading ? (
                <Skeleton className="w-20 h-4 flex">
                  <div className="w-full h-full bg-gray-300 rounded"></div>
                </Skeleton>
              ) : (
                <p className="text-gray-700">{pacienteData?.weight || 'N/A'} kg</p>
              )}
            </div>

            <div className='flex flex-col items-center'>
              <p className="text-gray-700 font-bold">Espécie</p>
              {isLoading ? (
                <Skeleton className="w-20 h-4 flex">
                  <div className="w-full h-full bg-gray-300 rounded"></div>
                </Skeleton>
              ) : (
                <p className="text-gray-700">{pacienteData?.species}</p>
              )}
            </div>
          </div>

          <div className='mt-4 sm:mt-0'>
            <Button className='flex' onClick={openModal}>Ver Paciente</Button>
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
