import React from 'react';
import { Button } from './components/ui/button';
import { Skeleton } from './components/ui/skeleton';

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

    const { consultation_date: date, observation: description, consultation_hour } = appointment;
  return (
    <div className="p-4 bg-white shadow-md rounded-lg border border-gray-300">
         <h2 className="text-lg font-semibold mb-2">Próximo Compromisso</h2>
        <div className='flex gap-20 items-center justify-center mt-5'>
       
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
     
      {/* patient */}
      <div className="mt-4">
        <h3 className="text-lg font-semibold mb-2">Paciente</h3>
        {/* card for patient */}
        <div className="flex items-center gap-4 p-4 bg-gray-100 rounded-lg">
          <div className="w-20 h-20 bg-gray-300 rounded-full">

          </div>
          <div className='flex gap-20 items-center text-center'>
            <div className='justify-center'>
            <p className="text-gray-700">Nome do Paciente</p>
            {/* skeleton */}
            <Skeleton className="w-20 h-4 flex">
                <div className="w-full h-full bg-gray-300 rounded"></div>
            </Skeleton>
            </div>
            
            <div className='justify-center'>
            <p className="text-gray-700">Raça</p>
            {/* skeleton */}
            <Skeleton className="w-20 h-4 flex">
                <div className="w-full h-full bg-gray-300 rounded"></div>
            </Skeleton>
            </div>

            <div className='justify-center'>
            <p className="text-gray-700">Idade</p>
            {/* skeleton */}
            <Skeleton className="w-20 h-4 flex">
                <div className="w-full h-full bg-gray-300 rounded"></div>
            </Skeleton>

            </div>

            <div className='justify-center'>
            <p className="text-gray-700">Peso</p>
            {/* skeleton */}
            <Skeleton className="w-20 h-4 flex">
                <div className="w-full h-full bg-gray-300 rounded"></div>
            </Skeleton>

            </div>

            <div className='justify-center'>
            <p className="text-gray-700">Espécie</p>
            {/* skeleton */}
            <Skeleton className="w-20 h-4 flex">
                <div className="w-full h-full bg-gray-300 rounded"></div>
            </Skeleton>

            </div>

            <Button className='flex'>Ver Paciente</Button>

          </div>
         
        </div>
      </div>
    </div>
  );
};

export default NextAppointment;
