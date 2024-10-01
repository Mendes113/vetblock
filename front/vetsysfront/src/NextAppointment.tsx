import React, { useState, useEffect } from 'react';
import { Button } from './components/ui/button';
import { Skeleton } from './components/ui/skeleton';
import PatientModal from './components/patientModal/patientmodal';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCalendarCheck } from '@fortawesome/free-solid-svg-icons';

interface Consultation {
  id: string; // UUID pode ser tratado como uma string
  reason: string;
  consultation_date: string;
  veterinary_crvm: string;
  consultation_hour: string;
  patient_id: string;
  observation: string;
}

const NextAppointment: React.FC<{ appointment: Consultation | null }> = ({ appointment }) => {
  if (!appointment) {
    return <div className="p-4 bg-gray-200 rounded-lg">Nenhum compromisso agendado.</div>;
  }

  const [isModalOpen, setIsModalOpen] = useState(false);
  const [pacienteData, setPacienteData] = useState<any | null>(null); // Dados do paciente
  const [isLoading, setIsLoading] = useState(true); // Estado de carregamento

  // Simulando a recuperação de dados do paciente assim que o componente for montado
  useEffect(() => {
    const fetchPacienteData = () => {
      setTimeout(() => {
        const fakePaciente = {
          name: "Rex",
          species: "Cachorro",
          breed: "Labrador",
          age: "5 anos",
          weight: "30kg",
          description: "Paciente amigável e brincalhão.",
          photoUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR1EVD_Qkrbtn2FvcBhkBHrbZ02q7vgdbTTpYc34CW23mCADFoQBKvKLbb7baIc8nqhF5s",
          lastConsultations: [
            { id: "1", date: "2021-09-01", reason: "Consulta de rotina" },
            { id: "2", date: "2021-08-01", reason: "Vermifugação" },
          ],
          lastMedications: ["Vermífugo 1x ao dia", "Antipulgas 1x ao mês"],
        };
        setPacienteData(fakePaciente);
        setIsLoading(false); // Finaliza o carregamento
      }, 1000); // Simula um atraso de 1 segundo
    };

    fetchPacienteData();
  }, []); // Executa apenas uma vez quando o componente é montado

  const openModal = () => {
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
  };

  const { consultation_date: date, observation: description, consultation_hour } = appointment;

  return (
    <div className="p-4 bg-white shadow-md rounded-lg border border-gray-300">
      <h2 className="text-lg font-semibold mb-2">Próximo Compromisso <FontAwesomeIcon icon={faCalendarCheck} /></h2>
      <div className="flex flex-col gap-4 items-center justify-center mt-5 sm:flex-row sm:gap-10">
        <p className="text-gray-700">
          <strong>Data:</strong> {new Date(date).toLocaleDateString()}
        </p>
        <p className="text-gray-700">
          <strong>Hora:</strong> {consultation_hour}
        </p>
        <p className="text-gray-700">
          <strong>Descrição:</strong> {description}
        </p>
      </div>

      {/* paciente */}
      <div className="mt-4">
        <h3 className="text-lg font-semibold mb-2">Paciente</h3>
        {/* card for patient */}
        <div className="flex flex-col items-center p-4 bg-gray-100 rounded-lg [w-1200]:flex-row sm:gap-4">
          {isLoading ? (
            <Skeleton className="w-24 h-24 rounded-full" />
          ) : (
            <img
              src={pacienteData?.photoUrl}
              alt={`Foto de ${pacienteData?.name}`}
              className="w-24 h-24 rounded-full object-cover"
            />
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
                <p className="text-gray-700">{pacienteData?.age}</p>
              )}
            </div>

            <div className='flex flex-col items-center'>
              <p className="text-gray-700 font-bold">Peso</p>
              {isLoading ? (
                <Skeleton className="w-20 h-4 flex">
                  <div className="w-full h-full bg-gray-300 rounded"></div>
                </Skeleton>
              ) : (
                <p className="text-gray-700">{pacienteData?.weight}</p>
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
