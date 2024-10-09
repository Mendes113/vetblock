import React, { useState } from 'react';
import { Card, CardContent } from '../components/ui/card';
import { Button } from '../components/ui/button'; // Botão existente

interface CreateConsultationProps {
  onSave: (consultation: NewConsultation) => void; // Função para salvar a nova consulta
}

interface NewConsultation {
  patientName: string;
  consultationDate: string;
  reason: string;
}

const CreateConsultation: React.FC<CreateConsultationProps> = ({ onSave }) => {
  const [patientName, setPatientName] = useState('');
  const [consultationDate, setConsultationDate] = useState('');
  const [reason, setReason] = useState('');

  const handleSave = () => {
    if (!patientName || !consultationDate || !reason) {
      alert('Preencha todos os campos!');
      return;
    }
    const newConsultation: NewConsultation = {
      patientName,
      consultationDate,
      reason
    };
    onSave(newConsultation);  // Enviando os dados para a função de salvamento
  };

  return (
    <div className="container mx-auto p-6">
      {/* Card simulando a criação da consulta como página */}
      <Card className="w-full max-w-lg mx-auto p-4 bg-white rounded-lg shadow-lg">
        <CardContent className="p-6">
          <h2 className="text-2xl font-bold mb-4 text-center">Agendar Nova Consulta</h2>

          {/* Formulário para a nova consulta */}
          <div className="mb-4">
            <label className="block text-gray-700 font-semibold mb-2" htmlFor="patientName">
              Nome do Paciente
            </label>
            <input
              id="patientName"
              type="text"
              className="w-full px-3 py-2 border rounded-lg"
              value={patientName}
              onChange={(e) => setPatientName(e.target.value)}
              placeholder="Digite o nome do paciente"
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700 font-semibold mb-2" htmlFor="consultationDate">
              Data da Consulta
            </label>
            <input
              id="consultationDate"
              type="date"
              className="w-full px-3 py-2 border rounded-lg"
              value={consultationDate}
              onChange={(e) => setConsultationDate(e.target.value)}
            />
          </div>

          <div className="mb-4">
            <label className="block text-gray-700 font-semibold mb-2" htmlFor="reason">
              Motivo da Consulta
            </label>
            <input
              id="reason"
              type="text"
              className="w-full px-3 py-2 border rounded-lg"
              value={reason}
              onChange={(e) => setReason(e.target.value)}
              placeholder="Descreva o motivo da consulta"
            />
          </div>

          {/* Botões para salvar */}
          <div className="flex justify-end gap-4">
            <Button className="bg-gray-400 text-white px-4 py-2 rounded-lg hover:bg-gray-500" onClick={() => window.history.back()}>
              Cancelar
            </Button>
            <Button className="bg-green-600 text-white px-4 py-2 rounded-lg hover:bg-green-700" onClick={handleSave}>
              Salvar Consulta
            </Button>
          </div>
        </CardContent>
      </Card>
    </div>
  );
};

export default CreateConsultation;
