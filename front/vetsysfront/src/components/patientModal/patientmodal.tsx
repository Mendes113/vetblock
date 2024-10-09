import React from "react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Skeleton } from "@/components/ui/skeleton";
import { Patient } from '../../Types';

type ModalPacienteProps = {
  patient: Patient | null;
  loading: boolean;
  onClose: () => void;
};

const InfoTable: React.FC<{ title: string; children: React.ReactNode }> = ({ title, children }) => (
  <div className="p-4 sm:p-2 bg-gray-50 rounded-lg shadow-sm w-full max-h-48 overflow-y-auto">
    <h4 className="mb-2 font-semibold text-base sm:text-sm text-gray-700">{title}</h4>
    <table className="w-full text-sm sm:text-xs text-left text-gray-700">
      <tbody>{children}</tbody>
    </table>
  </div>
);

const PatientModal: React.FC<ModalPacienteProps> = ({ patient, loading, onClose }) => {
  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-60 transition-opacity duration-300 ease-in-out">
      <div className="bg-white p-6 sm:p-4 rounded-lg shadow-2xl max-w-lg w-full sm:max-w-sm transition-transform transform scale-100 animate-fadeIn">
        <Card className="border-none">
          <CardHeader className="text-center">
            <CardTitle className="text-xl sm:text-lg text-gray-800">
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
                      alt={`Foto de ${patient.name || "paciente"}`}
                      className="w-24 h-24 sm:w-20 sm:h-20 rounded-full object-cover shadow-md"
                    />
                  ) : (
                    <Skeleton className="w-24 h-24 sm:w-20 sm:h-20 rounded-full" />
                  )}
                </div>
                <div className="text-sm sm:text-xs text-gray-700">
                  <p><strong>Espécie:</strong> {patient?.species || "Não disponível"}</p>
                  <p><strong>Raça:</strong> {patient?.breed || "Não disponível"}</p>
                  <p><strong>Idade:</strong> {patient?.age || "Não disponível"}</p>
                  <p><strong>Peso:</strong> {patient?.weight || "Não disponível"}</p>
                  <p><strong>Descrição:</strong> {patient?.description || "Não disponível"}</p>
                </div>

                <div className="mt-4">
                  <InfoTable title="Últimas Consultas">
                    {patient?.lastConsultations?.length ? (
                      patient.lastConsultations.map((consultation) => (
                        <tr key={consultation.consultation_id} className="border-b">
                          <td className="py-2">{new Date(consultation.consultation_date).toLocaleDateString()}</td>
                          <td className="py-2">{consultation.reason}</td>
                        </tr>
                      ))
                    ) : (
                      <tr>
                        <td colSpan={2} className="text-center py-2">
                          <Skeleton className="w-20 h-4 mx-auto" />
                        </td>
                      </tr>
                    )}
                  </InfoTable>
                  <div className="my-4" />
                  <InfoTable title="Últimas Medicações">
                    {patient?.lastMedications?.length ? (
                      patient.lastMedications.map((medication, index) => (
                        <tr key={index} className="border-b">
                          <td className="py-2">{medication}</td>
                        </tr>
                      ))
                    ) : (
                      <tr>
                        <td className="text-center py-2">
                          <Skeleton className="w-20 h-4 mx-auto" />
                        </td>
                      </tr>
                    )}
                  </InfoTable>
                </div>
              </> 
            )}
          </CardContent>
          <div className="mt-6 text-center pb-7">
            <Button className="w-full sm:w-auto text-sm sm:text-xs bg-red-600 hover:bg-red-700 text-white py-2 mx-auto" onClick={onClose}>
              Fechar
            </Button>
          </div>
        </Card>
      </div>
    </div>
  );
};

export default PatientModal;
