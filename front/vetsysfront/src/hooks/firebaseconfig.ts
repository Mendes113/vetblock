// src/hooks/firebaseconfig.ts
import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";
import { getAnalytics } from "firebase/analytics";

// Configuração do Firebase
const firebaseConfig = {
  apiKey: "AIzaSyDS2fnigTebl7y0VZm2EGH3AD8R3-1l-xk",
  authDomain: "vetsys-46e80.firebaseapp.com",
  projectId: "vetsys-46e80",
  storageBucket: "vetsys-46e80.appspot.com",
  messagingSenderId: "487118652609",
  appId: "1:487118652609:web:879926004fb9e382460040",
  measurementId: "G-8SQH8K6SY2"
};

// Inicializa o Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);

// Inicializa e exporta o serviço de autenticação
export const auth = getAuth(app);

export default app;
