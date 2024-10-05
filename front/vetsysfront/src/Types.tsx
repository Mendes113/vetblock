// types.ts
export interface Patient {
    id: string;
    name: string;
    species: string;
    breed: string;
    age: number;
    weight: number;
    photoUrl: string | null;
    description: string;
    lastConsultations: Consultation[];
    lastMedications: string[];
  }
  

  export interface Consultation {
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