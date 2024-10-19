import React, { useState } from "react";
import { signInWithEmailAndPassword, createUserWithEmailAndPassword, GoogleAuthProvider, signInWithPopup, getAuth } from "firebase/auth";
import axios from "axios"; // Adicionando axios para fazer requisições HTTP
import { useNavigate } from "react-router-dom";
import { Button } from "../components/ui/button";
import { Card } from "../components/ui/card";
import { Input } from "../components/ui/input";
import { auth } from "@/hooks/firebaseconfig";

const Login: React.FC = () => {
  const [email, setEmail] = useState<string>("");
  const [password, setPassword] = useState<string>("");
  const [cpf, setCpf] = useState<string>("");
  const [phone, setPhone] = useState<string>("");
  const [isSignUp, setIsSignUp] = useState<boolean>(false); // Estado para alternar entre login e registro
  const [isAnimalCardVisible, setIsAnimalCardVisible] = useState<boolean>(false); // Estado para controlar o card de adicionar animal
  const [animalName, setAnimalName] = useState<string>("");
  const [animalSpecies, setAnimalSpecies] = useState<string>("");
  const [animalAge, setAnimalAge] = useState<string>("");
  const navigate = useNavigate();


  const handleLogin = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const userCredential = await signInWithEmailAndPassword(auth, email, password);
      const user = userCredential.user;

      console.log("Login bem-sucedido para:", user.email);
      
      // Navega para a página inicial após o login
      navigate("/");
    } catch (error) {
      console.error("Erro ao fazer login:", error);
      alert("Erro ao fazer login. Verifique suas credenciais e tente novamente.");
    }
  };

  const handleSignUp = async (e: React.FormEvent) => {
    e.preventDefault();
    console.log("Sign Up Button Clicked"); 
    try {
      if (!email || !password || !cpf || !phone) {
        alert("Por favor, preencha todos os campos obrigatórios.");
        return;
      }
  
      // Envia uma requisição ao backend para criar o usuário
      await axios.post("http://localhost:8081/api/register", {
        email,
        password,
        cpf,
        phone,
        animal: {
          name: animalName,
          species: animalSpecies,
          age: animalAge
        }
      });
  
      console.log("Registro bem-sucedido");
      setIsAnimalCardVisible(true);
    } catch (error: any) {
      console.error("Erro ao fazer registro:", error);
      alert(`Erro ao fazer registro: ${error.message || error}`);
    }
  };
  

  const handleGoogleLogin = async () => {
    const provider = new GoogleAuthProvider();
    try {
      const result = await signInWithPopup(auth, provider);
      const user = result.user;
      
      console.log("Login com Google bem-sucedido para:", user.email);
      
      // Navega para a página inicial após o login
      navigate("/");
    } catch (error) {
      console.error("Erro ao fazer login com Google:", error);
      alert("Erro ao fazer login com Google. Tente novamente.");
    }
  };

  const handleAddAnimal = () => {
    console.log("Animal Adicionado:", {
      name: animalName,
      species: animalSpecies,
      age: animalAge
    });
    // Aqui você pode adicionar a lógica para salvar os dados do animal no banco de dados
    setIsAnimalCardVisible(false);
    navigate("/"); // Navega para a página inicial após adicionar o animal
  };

  return (
    <div className="flex h-screen bg-gradient-to-r from-blue-50 to-blue-100">
      {/* Seção de imagem à esquerda */}
      <div className="w-1/2 h-full hidden md:block">
        <img
          src="https://cdn.discordapp.com/attachments/1058806818224222278/1297207248896000040/DALLE_2024-10-19_11.38.06_-_Minimal_vector_illustration_of_a_veterinary_clinic_focused_on_dogs_and_cats._The_scene_features_a_veterinarian_gently_interacting_with_a_dog_and_a_cat.webp?ex=671515e5&is=6713c465&hm=ba41fb6e3b2a87a57584b6f054d54d7e551d019a71a12adb9bfdd6e729aa22d9&" // Substitua pelo caminho real da sua imagem
          alt="Ilustração"
          className="w-full h-full object-cover"
        />
      </div>

      {/* Card de login/registro à direita */}
      <div className="flex items-center justify-center w-full md:w-1/2 bg-white bg-opacity-70">
        {isAnimalCardVisible ? (
          <Card className="w-full max-w-md p-8 bg-white shadow-xl rounded-3xl">
            <h1 className="text-2xl font-semibold text-center mb-6 text-gray-800">
              Add Animal Details
            </h1>
            <Input
              type="text"
              placeholder="Animal Name"
              value={animalName}
              onChange={(e) => setAnimalName(e.target.value)}
              className="mb-4 border border-gray-300 rounded-lg px-4 py-2"
            />
            <Input
              type="text"
              placeholder="Species (e.g., Dog, Cat)"
              value={animalSpecies}
              onChange={(e) => setAnimalSpecies(e.target.value)}
              className="mb-4 border border-gray-300 rounded-lg px-4 py-2"
            />
            <Input
              type="text"
              placeholder="Age"
              value={animalAge}
              onChange={(e) => setAnimalAge(e.target.value)}
              className="mb-6 border border-gray-300 rounded-lg px-4 py-2"
            />
            <Button onClick={handleAddAnimal} className="w-full bg-blue-500 text-white py-2 rounded-full hover:bg-blue-600 transition duration-200">
              Add Animal
            </Button>
          </Card>
        ) : (
          <Card className="w-full max-w-md p-8 bg-white shadow-xl rounded-3xl">
            <h1 className="text-3xl font-semibold text-center mb-6 text-gray-800">
              {isSignUp ? "Create an Account" : "Sign In"}
            </h1>

            <form onSubmit={isSignUp ? handleSignUp : handleLogin}>
              <Input
                type="email"
                placeholder="name@example.com"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                className="mb-4 border border-gray-300 rounded-lg px-4 py-2"
              />
              <Input
                type="password"
                placeholder="Senha"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className="mb-4 border border-gray-300 rounded-lg px-4 py-2"
              />
              {isSignUp && (
                <>
                  <Input
                    type="text"
                    placeholder="CPF"
                    value={cpf}
                    onChange={(e) => setCpf(e.target.value)}
                    className="mb-4 border border-gray-300 rounded-lg px-4 py-2"
                  />
                  <Input
                    type="tel"
                    placeholder="Telefone"
                    value={phone}
                    onChange={(e) => setPhone(e.target.value)}
                    className="mb-6 border border-gray-300 rounded-lg px-4 py-2"
                  />
                </>
              )}
              <Button type="submit" variant="default" className="w-full bg-blue-500 text-white py-2 rounded-full hover:bg-blue-600 transition duration-200">
                {isSignUp ? "Sign Up with Email" : "Sign In with Email"}
              </Button>
            </form>

            {/* Linha separadora estilizada */}
            <div className="flex items-center my-6">
              <hr className="w-full border-gray-300" />
              <span className="px-2 text-gray-500">OR CONTINUE WITH</span>
              <hr className="w-full border-gray-300" />
            </div>

            <Button variant="outline" onClick={handleGoogleLogin} className="w-full border border-gray-300 py-2 rounded-full hover:bg-gray-200 transition duration-200">
              Continue with Google
            </Button>
            <Button variant="outline" className="w-full border border-gray-300 py-2 mt-2 rounded-full hover:bg-gray-200 transition duration-200">
              GitHub
            </Button>

            {/* Alternar entre Sign In e Sign Up */}
            <p className="text-center text-sm text-gray-600 mt-6">
              {isSignUp ? (
                <>
                  Already have an account?{" "}
                  <button onClick={() => setIsSignUp(false)} className="underline hover:text-blue-500">
                    Sign In
                  </button>
                </>
              ) : (
                <>
                  Don’t have an account?{" "}
                  <button onClick={() => setIsSignUp(true)} className="underline hover:text-blue-500">
                    Sign Up
                  </button>
                </>
              )}
            </p>

            {/* Termos e links adicionais */}
            <p className="text-center text-sm text-gray-600 mt-2">
              By clicking continue, you agree to our{" "}
              <a href="#" className="underline hover:text-blue-500">
                Terms of Service
              </a>{" "}
              and{" "}
              <a href="#" className="underline hover:text-blue-500">
                Privacy Policy
              </a>.
            </p>
          </Card>
        )}
      </div>
    </div>
  );
};

export default Login;
