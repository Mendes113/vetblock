import React, { useState } from 'react';
import Calendar from 'react-calendar';
import 'react-calendar/dist/Calendar.css'; // Inclua o CSS básico do react-calendar
import Modal from 'react-modal';

Modal.setAppElement('#root'); // Necessário para acessibilidade, #root é o id do elemento principal do React

export function Dashboard() {
  const [appointments, setAppointments] = useState([
    { date: new Date(2024, 10, 15), title: "Consulta de rotina" },
    { date: new Date(2024, 10, 20), title: "Exame de sangue" },
  ]);

  const [selectedDate, setSelectedDate] = useState<Date | null>(null);
  const [isModalOpen, setIsModalOpen] = useState(false); // Estado para controlar o modal

  // Função para abrir o modal e definir a data selecionada
  const handleDayClick = (value: Date) => {
    setSelectedDate(value);
    setIsModalOpen(true); // Abrir o modal ao clicar no dia
  };

  // Função para fechar o modal
  const closeModal = () => {
    setIsModalOpen(false);
  };

  const tileContent = ({ date, view }: { date: Date; view: string }) => {
    if (view === 'month') {
      const isAppointment = appointments.find((appt) => appt.date.toDateString() === date.toDateString());
      return isAppointment ? (
        <div className="flex justify-center items-center">
          <span className="w-3 h-3 bg-blue-500 rounded-full"></span>
        </div>
      ) : null;
    }
  };

  const selectedAppointments = appointments.filter(
    (appt) => selectedDate && appt.date.toDateString() === selectedDate.toDateString()
  );

  return (
    <div className="grid min-h-screen grid-cols-1 md:grid-cols-[220px_1fr] lg:grid-cols-[280px_1fr] xl:grid-cols-[300px_1fr] xl:max-w-[90%] 2xl:max-w-[80%] w-full mx-auto">
      {/* Sidebar e Header... */}

      <main className="flex flex-1 flex-col gap-4 p-4 sm:p-6 md:gap-6 lg:p-8 xl:gap-8 xl:p-10">
        <div className="flex items-center justify-between">
          <h1 className="text-xl font-semibold sm:text-2xl md:text-3xl xl:text-4xl">Appointments</h1>
        </div>

        {/* Calendário com consultas */}
        <div className="flex flex-1 flex-col items-center justify-center rounded-lg border border-dashed shadow-sm p-4 bg-white">
          <div className="w-full max-w-lg mb-6">
            <Calendar
              tileContent={tileContent}
              onClickDay={handleDayClick} // Selecionar o dia clicado
              className="p-4 border rounded-lg shadow-lg"
              tileClassName={({ date, view }) => {
                const hasAppointment = appointments.some((appt) => appt.date.toDateString() === date.toDateString());
                return hasAppointment ? 'bg-blue-100 text-blue-800 font-semibold rounded-full' : '';
              }}
            />
          </div>

          {/* Modal para exibir eventos do dia */}
          <Modal
            isOpen={isModalOpen}
            onRequestClose={closeModal}
            contentLabel="Modal de eventos do dia"
            className="modal"
            overlayClassName="modal-overlay"
          >
            <div className="p-4 bg-gray-100 rounded-lg shadow-md">
              <h2 className="text-lg font-semibold text-gray-800 mb-2">
                Eventos para {selectedDate?.toDateString()}:
              </h2>
              {selectedAppointments.length > 0 ? (
                <ul>
                  {selectedAppointments.map((appt, index) => (
                    <li key={index} className="text-gray-700">
                      - {appt.title}
                    </li>
                  ))}
                </ul>
              ) : (
                <p className="text-gray-500">Nenhum evento para este dia.</p>
              )}
              <button
                onClick={closeModal}
                className="mt-4 px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-700"
              >
                Fechar
              </button>
            </div>
          </Modal>

          {/* Legenda explicativa */}
          <div className="flex items-center gap-2 mt-4">
            <span className="w-3 h-3 bg-blue-500 rounded-full"></span>
            <span className="text-sm text-gray-700">Consulta marcada</span>
          </div>
        </div>
      </main>
    </div>
  );
}
