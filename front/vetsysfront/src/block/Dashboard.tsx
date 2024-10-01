import Link from "next/link";
import {
  Bell,
  CircleUser,
  Home,
  LineChart,
  Menu,
  Package,
  Package2,
  Search,
  ShoppingCart,
  Users,
} from "lucide-react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faDog, faNotesMedical, faCalendar } from "@fortawesome/free-solid-svg-icons";
import React, { useState, useEffect } from "react";
import { Calendar } from "react-calendar";
import "react-calendar/dist/Calendar.css"; // CSS do react-calendar

import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Input } from "@/components/ui/input";
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet";
import NextAppointment from "@/NextAppointment";
import {Navbar} from "@/components/Nabar/navbar";
import TopBar from "@/components/topbar";

export const description =
  "A products dashboard with a sidebar navigation and a main content area. The dashboard has a header with a search input and a user menu. The sidebar has a logo, navigation links, and a card with a call to action. The main content area shows an empty state with a call to action.";

export function Dashboard() {
  const [date, setDate] = useState(new Date());
  const [consultations, setConsultations] = useState<Consultation[]>([]);
  const [appointment, setAppointment] = useState<Consultation | null>(null);

  // Fake appointment data
  const fakeAppointment = {
    id: "1",
    reason: "Consulta de rotina",
    consultation_date: "2022-12-31",
    veterinary_crvm: "123456",
    consultation_hour: "08:00",
    patient_id: "1",
    observation: "Observação de teste",
  };

  useEffect(() => {
    const fetchConsultations = async () => {
      // Mock data, could be fetched from an API
      const data = [fakeAppointment];
      setConsultations(data);
    };

    fetchConsultations();
  }, []);

  type Consultation = {
    id: string;
    reason: string;
    consultation_date: string;
    veterinary_crvm: string;
    consultation_hour: string;
    patient_id: string;
    observation: string;
  };

  return (
    <div className="min-h-screen flex flex-col">
      {/* TopBar stays fixed at the top */}
      {/* <TopBar /> */}

      <div className="flex flex-1">
        {/* Sidebar for Navbar */}
        <aside className="">
          <Navbar />
        </aside>

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
              <NextAppointment appointment={fakeAppointment} />
            </div>
          </div>
        </main>
      </div>
    </div>
  );
}
