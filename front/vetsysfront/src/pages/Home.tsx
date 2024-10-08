
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faDog, faNotesMedical, faCalendar } from "@fortawesome/free-solid-svg-icons";
import React, { useState, useEffect } from "react";
import { Calendar } from "react-calendar";
import "react-calendar/dist/Calendar.css"; // CSS do react-calendar


import NextAppointment from "@/NextAppointment";
import TopBar from "@/components/topbar";
import {Navbar} from "@/components/Navbar/navbar";
import Hospitalizations from "@/components/hospitalizations/hospitalizations";
import Consultations from "@/components/consultations/consultations";
import { Footer } from "@/components/Footer/Footer";

export const description =
  "A products dashboard with a sidebar navigation and a main content area. The dashboard has a header with a search input and a user menu. The sidebar has a logo, navigation links, and a card with a call to action. The main content area shows an empty state with a call to action.";

  export function Home() {
    const [date, setDate] = useState(new Date());
    const [consultations, setConsultations] = useState<Consultation[]>([]);
    const [appointment, setAppointment] = useState<Consultation | null>(null);
    const crvm = "56789"; // Você pode alterar este valor ou passá-lo dinamicamente
  
    // Definição do tipo Consultation
    interface Patient {
      animal_id: string;
      id: string; // ID do paciente
      name: string;
      species: string;
      breed: string;
      age: number;
      weight: number;
      image: string | null; // URL da foto do paciente (pode ser nulo)
      description: string;
      photoUrl: string | null; // URL da foto do paciente (opcional)
      lastConsultations: Consultation[]; // Últimas consultas
      lastMedications: string[]; // Últimas medicações
    }
    
    interface Consultation {
      consultation_id: string;
      animal_id: string;
      crvm: string;
      consultation_date: string;
      consultation_hour: string;
      observation: string;
      reason: string;
      consultation_type: string;
      consultation_description: string;
      consultation_prescription: string;
      consultation_price: number;
      consultation_status: string;
    }
  
    useEffect(() => {
      const fetchConsultations = async () => {
        try {
          const response = await fetch(`http://localhost:8081/api/v1/consultations/${crvm}`);
          if (!response.ok) {
            throw new Error("Erro ao buscar as consultas");
          }
          const dataConsult = await response.json();
          console.log(dataConsult);
          setConsultations(dataConsult);
        } catch (error) {
          console.error("Erro ao buscar dados:", error);
        }
      };

      const fetchNextAppointment = async () => {
        try {
          const response = await fetch(`http://localhost:8081/api/v1/veterinary/${crvm}/next-consultation`);
          if (!response.ok) {
            console.log(response);
            throw new Error("Erro ao buscar o próximo compromisso");
          }
          const dataAppointment = await response.json();
          console.log(dataAppointment);
          setAppointment(dataAppointment);
        } catch (error) {
          console.error("Erro ao buscar dados:", error);
        }
      }

      
      // fetchAnimal("12345");
      fetchNextAppointment();
      fetchConsultations();
    }, [crvm]); // 'crvm' adicionado como dependência para garantir que a consulta aconteça se o valor de 'crvm' mudar
  
   
  return (
    <div className="min-h-screen flex flex-col">
      {/* TopBar stays fixed at the top */}
      {/* <TopBar /> */}

      <div className="flex flex-1">
        {/* Sidebar for Navbar */}
      

        {/* Main content area */}
        <main className="flex-1 p-4">
          <div className="flex flex-col gap-4">
            {/* Calendar */}
            <div className="w-full h-full">
              <h2 className="text-xl font-semibold">
                Calendário <FontAwesomeIcon icon={faCalendar} className="ml-1" />
              </h2>

              <Calendar
                value={date}
                className="mt-2 border p-4 rounded-lg max-h-full max-w-full w-full h-full"
              />
            </div>

            {/* Next Appointment */}
            <div>
              <NextAppointment appointment={appointment} />
            </div>

            <div className="flex flex-row gap-16 align-middle justify-center">
            <div className="flex-1 max-h-max">
              <Hospitalizations />
            </div>
            <div className="flex-1">
              <Consultations />
            </div>
          
              </div>
            
          </div>
        </main>
      </div>
     
    </div>
  );
}
