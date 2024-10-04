import React, { useState, useEffect } from 'react';
import { Button } from '../ui/button';
import { Skeleton } from '../ui/skeleton';
import PatientModal from '../patientModal/patientmodal';
import { Carousel, CarouselContent, CarouselItem, CarouselNext, CarouselPrevious } from '../ui/carousel';
import { Card, CardContent } from '../ui/card';
import { Chart } from '../char/chart';
import { AddMedicationCard } from '../addmed/addMedication';

interface Patient {
  id: string;
  name: string;
  species: string;
  breed: string;
  age: string;
  weight: string;
  description: string;
  photoUrl: string;
  lastConsultations: { id: string; date: string; reason: string }[];
  lastMedications: string[];
}

const Hospitalizations: React.FC = () => {
  const [patients, setPatients] = useState<Patient[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [selectedPatient, setSelectedPatient] = useState<Patient | null>(null);
  const [isAddMedicationModalOpen, setIsAddMedicationModalOpen] = useState(false);

  interface Dosage {
    medicationId: string;
    startDate: Date;
    endDate: Date;
    quantity: number;
    dosage: string;
    administrationRoute: string;
    frequency: string;
    durationDays: number;

  }

  interface NewMedication {
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
  
  const handleAddMedication = (medication: NewMedication, dosage: Dosage) => {
    // Aqui você pode adicionar lógica para enviar esses dados para o backend
    console.log('Medicação adicionada:', medication);
    console.log('Dosagem adicionada:', dosage);
  };

  // Simulando a recuperação de dados de pacientes internados
  useEffect(() => {
    const fetchPatients = () => {
      setTimeout(() => {
        const fakePatients = [
          {
            id: "1",
            name: "Rex",
            species: "Cachorro",
            breed: "Labrador",
            age: "5 anos",
            weight: "30kg",
            description: "Paciente amigável e brincalhão.",
            lastConsultations: [
              { id: "1", date: "2021-09-01", reason: "Consulta de rotina" },
              { id: "2", date: "2021-08-01", reason: "Vermifugação" },
            ],
            lastMedications: ["Vermífugo 1x ao dia", "Antipulgas 1x ao mês"],
            photoUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR1EVD_Qkrbtn2FvcBhkBHrbZ02q7vgdbTTpYc34CW23mCADFoQBKvKLbb7baIc8nqhF5s",
          },
          {
            id: "2",
            name: "Bobby",
            species: "Cachorro",
            breed: "Beagle",
            age: "3 anos",
            weight: "12kg",
            description: "Paciente enérgico e curioso.",
            lastConsultations: [
              { id: "1", date: "2021-09-01", reason: "Consulta de rotina" },
              { id: "2", date: "2021-08-01", reason: "Vermifugação" },
            ],
            lastMedications: ["Vermífugo 1x ao dia", "Antipulgas 1x ao mês"],
            photoUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRYLhgYygqgkhWTAZniKnHXIZaYccVdcOwn4w&s",
          },
          {
            id: "3",
            name: "Felix",
            species: "Gato",
            breed: "Siamês",
            age: "2 anos",
            weight: "5kg",
            description: "Paciente calmo e independente.",
            lastConsultations: [
              { id: "1", date: "2021-09-01", reason: "Consulta de rotina" },
              { id: "2", date: "2021-08-01", reason: "Vermifugação" },
            ],
            lastMedications: ["Vermífugo 1x ao dia", "Antipulgas 1x ao mês"],
            photoUrl: "data:image/jpeg;base64,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",
          },
        ];
        setPatients(fakePatients);
        setIsLoading(false);
      }, 1000); // Simulando atraso de 1 segundo
    };

    fetchPatients();
  }, []);

  const openModal = (patient: Patient) => {
    setSelectedPatient(patient);
    setIsModalOpen(true);
  };

  const closeModal = () => {
    setIsModalOpen(false);
    setSelectedPatient(null);
  };

  const openModaMedication = (patient: Patient) => {
    setIsAddMedicationModalOpen(true);
    setSelectedPatient(patient);
  }


return (
  <div className="p-6 bg-gray-50 shadow-lg rounded-xl border border-gray-200 items-center justify-center mx-auto">
    <h2 className="text-xl font-semibold mb-6 text-gray-800">Pacientes Internados</h2>
    <div className='ml-0 mr-0 flex justify-center pb-2'>
      <Chart />
    </div>
    
    <Carousel className="w-full max-w-3xl mx-auto">
      <CarouselContent className="flex -ml-1">
        {patients.map((patient) => (
          <CarouselItem key={patient.id} className="pl-1 w-1/3">
            <div className="p-2">
              <Card className="transition-shadow hover:shadow-md rounded-lg">
                <CardContent className="p-6 bg-white rounded-lg shadow-sm">
                  {isLoading ? (
                    <Skeleton className="w-24 h-24 rounded-full" />
                  ) : (
                    <img
                      src={patient.photoUrl}
                      alt={`Foto de ${patient.name}`}
                      className="w-24 h-24 rounded-full object-cover mx-auto transition-transform hover:scale-105"
                    />
                  )}
                  <div className="text-center mt-4">
                    <h3 className="text-lg font-bold text-gray-900">{patient.name}</h3>
                    <p className="text-sm text-gray-600">Idade: {patient.age}</p>
                    <p className="text-sm text-gray-600">Tempo Internado: 2 dias</p>
                    <p className="text-sm text-gray-600">Última Medicação: {patient.lastMedications[0]}</p>
                    <p className="text-sm text-gray-600">Próxima Medicação: 12h</p>
                    <div className='gap-2'>
                      <Button className="mt-4 bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors mr-2" onClick={() => openModal(patient)}>
                        Ver Detalhes
                      </Button>
                      <Button className="mt-4 bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition-colors" onClick={() => openModaMedication(patient)}>
                        Adicionar Medicação
                      </Button>
                    </div>
                  </div>
                </CardContent>
              </Card>
            </div>
          </CarouselItem>
        ))}
      </CarouselContent>
      <CarouselPrevious />
      <CarouselNext />
    </Carousel>

    {isModalOpen && selectedPatient && (
      <PatientModal patient={selectedPatient} loading={isLoading} onClose={closeModal} />
    )}

    {isAddMedicationModalOpen && selectedPatient && (
      <AddMedicationCard
        onClose={() => setIsAddMedicationModalOpen(false)}
        onAddMedication={handleAddMedication}
        patient={selectedPatient}
      />
    )}
  </div>
);

};

export default Hospitalizations;
