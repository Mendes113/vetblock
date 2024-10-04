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
      <div className="bg-white p-4 sm:p-2 rounded-lg shadow-lg max-w-md w-full sm:max-w-sm ">
        <Card>
          <CardHeader>
            <CardTitle className="text-lg sm:text-base">
              Paciente: {loading ? <Skeleton className="w-32" /> : patient?.name || "Paciente Desconhecido"}
            </CardTitle>
          </CardHeader>
          <CardContent>
            {loading ? (
              <>
                <div className="flex justify-center mb-4">
                  <Skeleton className="w-24 h-24 sm:w-20 sm:h-20 rounded-full" />
                </div>
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
                      className="w-24 h-24 sm:w-20 sm:h-20 rounded-full object-cover"
                    />
                  ) : (
                    <Skeleton className="w-24 h-24 sm:w-20 sm:h-20 rounded-full" />
                  )}
                </div>
                <div className="text-sm sm:text-xs">
                  <p><strong>Espécie:</strong> {patient?.species || "Não disponível"}</p>
                  <p><strong>Raça:</strong> {patient?.breed || "Não disponível"}</p>
                  <p><strong>Idade:</strong> {patient?.age || "Não disponível"}</p>
                  <p><strong>Peso:</strong> {patient?.weight || "Não disponível"}</p>
                  <p><strong>Descrição:</strong> {patient?.description || "Não disponível"}</p>
                </div>
                <div className="mt-2 p-1">
                  <div className="flex items-center gap-4 p-4 sm:p-2 bg-gray-100 rounded-lg flex-col">
                    <h4 className="mt-4 font-semibold sm:mt-2">Últimas Consultas:</h4>
                    <ul className="list-disc ml-5 sm:ml-3">
                      {patient && patient.lastConsultations && patient.lastConsultations.length ? (
                        <>
                          {patient.lastConsultations.map((consultation) => (
                            <li key={consultation.id} className="text-sm sm:text-xs">
                              {new Date(consultation.date).toLocaleDateString()}: {consultation.reason}
                            </li>
                          ))}
                          <Button className="mt-2 sm:mt-1 text-sm sm:text-xs" onClick={onClose}>Acessar Consultas</Button>
                        </>
                      ) : (
                        <Skeleton className="w-20 h-4" />
                      )}
                    </ul>
                  </div>
                  <div className="my-2" />
                  <div className="flex items-center gap-4 p-4 sm:p-2 bg-gray-100 rounded-lg flex-col">
                    <h4 className="mt-4 font-semibold sm:mt-2">Últimas Medicações:</h4>
                    <ul className="list-disc ml-5 sm:ml-3">
                      {patient && patient.lastMedications && patient.lastMedications.length ? (
                        <>
                          {patient.lastMedications.map((medication, index) => (
                            <li key={index} className="text-sm sm:text-xs">{medication}</li>
                          ))}
                          <Button className="mt-2 sm:mt-1 text-sm sm:text-xs" onClick={onClose}>Acessar Medicações</Button>
                        </>
                      ) : (
                        <Skeleton className="w-20 h-4" />
                      )}
                    </ul>
                  </div>
                </div>
              </>
            )}
          </CardContent>
          <div className="mt-4 pb-2">
            <Button className="w-full sm:w-auto text-sm sm:text-xs" onClick={onClose}>Fechar</Button>
          </div>
        </Card>
      </div>
    </div>
  );
};

export default PatientModal;
