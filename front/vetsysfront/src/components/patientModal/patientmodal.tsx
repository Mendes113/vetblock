import React from "react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton"; // Importe o componente Skeleton

type Patient = {
  id: string;
  name: string;
  species: string;
  breed: string;
  age: string;
  weight: string;
  description: string;
  photoUrl: string; // URL da foto do paciente
  lastConsultations: Consultation[]; // Últimas consultas
  lastMedications: string[]; // Últimas medicações
};

type Consultation = {
  id: string;
  date: string; // Data da consulta
  reason: string; // Motivo da consulta
};

type ModalPacienteProps = {
  patient: Patient | null;
  loading: boolean; // Adicione a propriedade de carregamento
  onClose: () => void;
};

const PatientModal: React.FC<ModalPacienteProps> = ({ patient, loading, onClose }) => {
  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
      <div className="bg-white p-4 rounded-lg shadow-lg max-w-md w-full">
        <Card>
          <CardHeader>
            <CardTitle>
              Paciente: {loading ? <Skeleton className="w-32" /> : patient?.name || "Paciente Desconhecido"}
            </CardTitle>
          </CardHeader>
          <CardContent>
            {loading ? (
              <>
                <div className="flex justify-center mb-4">
                  <Skeleton className="w-24 h-24 rounded-full" />
                </div>
                <Skeleton className="h-4 w-3/4 mb-2" />
                <Skeleton className="h-4 w-3/4 mb-2" />
                <Skeleton className="h-4 w-3/4 mb-2" />
                <Skeleton className="h-4 w-3/4 mb-2" />
              </>
            ) : (
              <>
                <div className="flex justify-center mb-4">
                  {patient?.photoUrl ? (
                    <img
                      src={patient.photoUrl}
                      alt={`Foto de ${patient.name}`}
                      className="w-24 h-24 rounded-full object-cover"
                    />
                  ) : (
                    <Skeleton className="w-24 h-24 rounded-full" />
                  )}
                </div>
                <div className="">

             
                <p><strong>Espécie:</strong> {patient?.species || "Não disponível"}</p>
                <p><strong>Raça:</strong> {patient?.breed || "Não disponível"}</p>
                <p><strong>Idade:</strong> {patient?.age || "Não disponível"}</p>
                <p><strong>Peso:</strong> {patient?.weight || "Não disponível"}</p>
                <p><strong>Descrição:</strong> {patient?.description || "Não disponível"}</p>

                </div>
                <div className="mt-2 p-1">
                <div className="flex items-center gap-4 p-4 bg-gray-100 rounded-lg flex-col">
                <h4 className="mt-4 font-semibold">Últimas Consultas:</h4>
                <ul className="list-disc ml-5">
                  {patient && patient.lastConsultations && patient.lastConsultations.length ? (
                    <>
                      {patient.lastConsultations.map((consultation) => (
                        <li key={consultation.id}>
                          {new Date(consultation.date).toLocaleDateString()}: {consultation.reason}
                        </li> 
                      ))}
                      <Button className="mt-2" onClick={onClose}>Acessar Consultas</Button>
                    </>
                  ) : (
                    <Skeleton className="w-20 h-4 flex">
                     <div className="w-full h-full bg-gray-300 rounded"></div>
                    </Skeleton>
                  )}

                        
                </ul>
                </div>
                <div className="my-2" />
                <div className="flex items-center gap-4 p-4 bg-gray-100 rounded-lg flex-col">
                <h4 className="mt-4 font-semibold">Últimas Medicações:</h4>
                <ul className="list-disc ml-5">
                  {patient && patient.lastMedications && patient.lastMedications.length ? (
                     <>
                    {patient.lastMedications.map((medication, index) => (
                       
                         <li key={index}>{medication}</li>
                        
                        
                     
                    ))}
                        <Button className="mt-2" onClick={onClose}>Acessar Medicações</Button>

                    </>
                  ) : (
                    <Skeleton className="w-20 h-4 flex">
                    <div className="w-full h-full bg-gray-300 rounded"></div>
                   </Skeleton>
                  )}
                </ul>
                </div>
                </div>
               
              </>
            )}
          </CardContent>
          <div className="mt-4 pb-2">
            <Button onClick={onClose}>Fechar</Button>
          </div>
        </Card>
      </div>
    </div>
  );
};

export default PatientModal;
