import React, { useState, useEffect } from 'react';
import { Button } from '../ui/button';
import { Skeleton } from '../ui/skeleton';
import PatientModal from '../patientModal/patientmodal';
import { Carousel, CarouselContent, CarouselItem, CarouselNext, CarouselPrevious } from '../ui/carousel';
import { Card, CardContent } from '../ui/card';
import { Chart } from '../char/chart';
import { AddMedicationCard } from '../addmed/addMedication';
import { ChartPie } from '../char/chartpie';

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

const Consultations: React.FC = () => {
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
    console.log('Medicação adicionada:', medication);
    console.log('Dosagem adicionada:', dosage);
  };

  // Simulando a recuperação de dados de pacientes
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
              { id: "1", date: "2023-10-01", reason: "Consulta de rotina" },
              { id: "2", date: "2023-09-01", reason: "Vacinação" },
            ],
            lastMedications: ["Vermífugo 1x ao dia", "Antipulgas 1x ao mês"],
            photoUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRYLhgYygqgkhWTAZniKnHXIZaYccVdcOwn4w&s"
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
              { id: "1", date: "2023-10-05", reason: "Consulta de rotina" },
              { id: "2", date: "2023-08-15", reason: "Exame de sangue" },
            ],
            lastMedications: ["Antibiótico 2x ao dia", "Vermífugo 1x ao dia"],
            photoUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRYLhgYygqgkhWTAZniKnHXIZaYccVdcOwn4w&s"
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
              { id: "1", date: "2023-09-25", reason: "Check-up" },
              { id: "2", date: "2023-08-20", reason: "Vacinação" },
            ],
            lastMedications: ["Antipulgas 1x ao mês"],
            photoUrl: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRYLhgYygqgkhWTAZniKnHXIZaYccVdcOwn4w&s"
          },
        ];
        setPatients(fakePatients);
        setIsLoading(false);
      }, 1000);
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

  const openModalMedication = (patient: Patient) => {
    setIsAddMedicationModalOpen(true);
    setSelectedPatient(patient);
  };

  return (
    <div className="p-6 bg-gray-50 shadow-lg rounded-xl border border-gray-200 mx-auto max-w-7xl">
      <h2 className="text-2xl font-bold mb-8 text-gray-800 text-center">Próximas Consultas</h2>

      {/* Gráfico para mostrar o histórico das consultas */}
      <div className="mb-8 flex justify-center">
        <ChartPie />
      </div>

      {/* Carrossel mostrando informações dos pacientes */}
      <Carousel className="w-full max-w-5xl mx-auto">
        <CarouselContent className="flex -ml-2 w-[500px]">
          {patients.map((patient) => (
            <CarouselItem key={patient.id} className="pl-2 w-full sm:w-1/2">
              <div className="p-3">
                <Card className="transition-all hover:shadow-xl rounded-lg hover:-translate-y-2">
                  <CardContent className="p-6 bg-white rounded-lg shadow-sm">
                    {/* Foto do Paciente */}
                    {isLoading ? (
                      <Skeleton className="w-20 h-20 rounded-full" />
                    ) : (
                      <img
                        src={patient.photoUrl}
                        alt={`Foto de ${patient.name}`}
                        className="w-20 h-20 rounded-full object-cover mx-auto transition-transform duration-300 hover:scale-105"
                      />
                    )}

                    {/* Informações do Paciente */}
                    <div className="text-center mt-4">
                      <h3 className="text-lg font-bold text-gray-800">{patient.name}</h3>
                      <p className="text-sm text-gray-600">Idade: {patient.age}</p>
                      <p className="text-sm text-gray-600">Última Consulta: {patient.lastConsultations[0]?.date}</p>
                      <p className="text-sm text-gray-600">Próxima Medicação: {patient.lastMedications[0]}</p>

                      {/* Botões de Ação */}
                      <div className="flex justify-center gap-4 mt-4">
                        <Button 
                          className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition-colors" 
                          onClick={() => openModal(patient)}
                        >
                          Ver Detalhes
                        </Button>
                        <Button 
                          className="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700 transition-colors" 
                          onClick={() => openModalMedication(patient)}
                        >
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

        <CarouselPrevious className="text-gray-500 hover:text-gray-700" />
        <CarouselNext className="text-gray-500 hover:text-gray-700" />
      </Carousel>

      {/* Modais */}
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

export default Consultations;
