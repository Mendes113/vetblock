import React, { useState, useEffect } from "react";
import axios from "axios";
import { Card, CardContent, CardTitle, CardDescription } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faInfoCircle, faDog, faCat } from "@fortawesome/free-solid-svg-icons";

// Componentes de Carrossel (presumindo que já existem no seu projeto)
import { Carousel, CarouselContent, CarouselItem, CarouselNext, CarouselPrevious } from "@/components/ui/carousel";

interface User {
  id: string;
  name: string;
  email: string;
  photoUrl: string;
}

interface Pet {
  id: string;
  name: string;
  species: string;
  breed: string;
  age: number;
  image: string;
}

interface Consultation {
  id: string;
  date: string;
  reason: string;
  value: number;
}

const UserPage: React.FC = () => {
  const [user, setUser] = useState<User | null>(null);
  const [pets, setPets] = useState<Pet[]>([]);
  const [consultations, setConsultations] = useState<Consultation[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchUserData();
    fetchUserPets();
    fetchUserConsultations();
  }, []);

  const fetchUserData = async () => {
    try {
      const response = await axios.get<User>("http://localhost:8081/api/v1/user");
      setUser(response.data);
      console.log("User data fetched:", response.data);
    } catch (error) {
      console.error("Erro ao buscar dados do usuário:", error);
      // Dados mock para teste:
      setUser({
        id: "1",
        name: "João da Silva",
        email: "joao@gmail.com",
        photoUrl: "https://randomuser.me/api/portraits/men/1.jpg", // Exemplo de URL válida
      });
    }
  };

  const fetchUserPets = async () => {
    try {
      const response = await axios.get<Pet[]>("http://localhost:8081/api/v1/user/pets");
      setPets(response.data);
      console.log("Pets fetched:", response.data);
    } catch (error) {
      console.error("Erro ao buscar pets do usuário:", error);
      // Dados mock para teste:
      setPets([
        {
          id: "1",
          name: "Rex",
          species: "Cachorro",
          breed: "Labrador",
          age: 3,
          image: "https://tenor.com/view/sus-suspicious-cat-suspicious-cat-gif-9601770602284137045.gif", // Exemplo de GIF
        },
        {
          id: "2",
          name: "Mia",
          species: "Gato",
          breed: "Siamês",
          age: 2,
          image: "https://tenor.com/view/sus-cat-2-suspicious-cat-the-cat-looks-suspiciously-cat-sits-in-front-of-food-the-ginger-cat-is-watching-gif-14890167989997543813.gif", // Exemplo de GIF
        },
        {
          id: "3",
          name: "Luna",
          species: "Gato",
          breed: "Persa",
          age: 4,
          image: "https://cdn.discordapp.com/attachments/1058806818224222278/1296998513137291274/9k.png?ex=6714537e&is=671301fe&hm=05b82ba2b00e97aa527ec90d6c25e286bab3f0862ccdae12e767464012f6d8ef&", // Imagem estática
        },
        {
          id: "4",
          name: "Osvaldo",
          species: "Gato",
          breed: "Persa",
          age: 4,
          image: "https://tenor.com/view/cat-chaterale-mdrr-3d-gif-17442785.gif",
        },
      ]);
    }
  };

  const fetchUserConsultations = async () => {
    try {
      const response = await axios.get<Consultation[]>("http://localhost:8081/api/v1/user/consultations");
      setConsultations(response.data);
      console.log("Consultations fetched:", response.data);
    } catch (error) {
      console.error("Erro ao buscar consultas do usuário:", error);
      // Dados mock para teste:
      setConsultations([
        {
          id: "1",
          date: "2023-10-12",
          reason: "Vacinação",
          value: 120.0,
        },
      ]);
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return <div className="text-center p-6">Carregando...</div>;
  }



  const RomanceCarousel: React.FC = () => (
    <Carousel className="w-full max-w-3xl mx-auto my-6">
      <CarouselContent className="flex">
        <CarouselItem className="flex-none w-full px-2">
        <Card className="bg-pink-100 rounded-xl shadow-md hover:shadow-lg transition-shadow duration-300">
  <img 
    src="https://tenor.com/view/spider-man-tobey-maguire-sad-cry-crying-gif-4535637.gif" // Substitua pelo caminho da sua imagem
    alt="Desde que te conheci..."
    className="w-full object-contain rounded-t-xl"
  />
  <CardContent>
    <h3 className="text-lg font-semibold text-center mb-2">Quando te conheci</h3>
    <p className="text-gray-600 text-center">Eu tava meio meh</p>
    <p className="text-gray-600 text-center">Achando que tava faltando algo</p>



  </CardContent>
</Card>

        </CarouselItem>
        
        <CarouselItem className="flex-none w-full px-2">
                    <Card className="bg-pink-100 rounded-xl shadow-md hover:shadow-lg transition-shadow duration-300">
            <img 
                src="https://cdn.discordapp.com/attachments/1058806818224222278/1297248861387227206/efea0b5b-7b6e-4169-91bc-0c800d4a4926.jpeg?ex=67153ca6&is=6713eb26&hm=a6d038dc6a6ef8c53c0b11f4f2f1081df6303eeb444294f5cfb004ca854fb420&" // Substitua pelo caminho da sua imagem
                alt="Desde que te conheci..."
                className="w-80 object-contain rounded-t-xl max-h-lvh mx-auto"
            />
            <CardContent>
                <h3 className="text-lg font-semibold text-center mb-2">Talvez</h3>
                <p className="text-gray-600 text-center">Acho que faltavam gatos na minha vida</p>
                <p className="text-gray-600 text-center">E ai você apareceu,</p>
                <p className="text-gray-600 text-center">Bom, na verdade...</p>




            </CardContent>
            </Card>
        </CarouselItem>
  
        <CarouselItem className="flex-none w-full px-2">
        <Card className="bg-pink-100 rounded-xl shadow-md hover:shadow-lg transition-shadow duration-300">
            <img 
                src="https://s3.ezgif.com/tmp/ezgif-3-5a72e309e6.gif" // Substitua pelo caminho da sua imagem
                alt="Desde que te conheci..."
                className="w-80 object-contain rounded-t-xl max-h-lvh flex mx-auto"
            />
            <CardContent>
                <h3 className="text-lg font-semibold text-center mb-2">Eu apareci</h3>
                <p className="text-gray-600 text-center">Com esse babaquinha</p>
                <p className="text-gray-600 text-center">Não tenho o que dizer, realmente queria muito ver você</p>




            </CardContent>
            </Card>
        </CarouselItem>
        <CarouselItem className="flex-none w-full px-2">
        <Card className="bg-pink-100 rounded-xl shadow-md hover:shadow-lg transition-shadow duration-300">
            <img 
                src="https://i.pinimg.com/originals/23/a9/4a/23a94a6623bce3ad6df176c0f997833d.jpg" // Substitua pelo caminho da sua imagem
                alt="Desde que te conheci..."
                className="w-full object-contain rounded-t-xl max-h-lvh flex mx-auto"
            />
            <CardContent>
                <h3 className="text-lg font-semibold text-center mb-2">Acho que chegou a hora</h3>
                <p className="text-gray-600 text-center">Se quiser correr é agora</p>

            </CardContent>
            </Card>
        </CarouselItem>
        <CarouselItem className="flex-none w-full px-2">
        <Card className="bg-pink-100 rounded-xl shadow-md hover:shadow-lg transition-shadow duration-300">
            <img 
                src="https://media.discordapp.net/attachments/1058805305112285258/1297253924813799474/ezgif-2-6e269c04d3.png?ex=6715415d&is=6713efdd&hm=7997aabffd884e4962caf5e73539984ba8cc222dac3c127c1a9c7d726a993d47&=&format=webp&quality=lossless&width=261&height=468" // Substitua pelo caminho da sua imagem
                alt="Desde que te conheci..."
                className="w-80 object-contain rounded-t-xl max-h-lvh flex mx-auto"
            />
            <CardContent>
                <h3 className="text-lg font-semibold text-center mb-2">Caso Contrario</h3>
                <p className="text-gray-600 text-center">Eu queria saber se teria a humildade de me deixar viver um pouco mais com o tigrinho e a sequela</p>

            </CardContent>
            </Card>
        </CarouselItem>
        <CarouselItem className="flex-none w-full px-2">
        <Card className="bg-pink-100 rounded-xl shadow-md hover:shadow-lg transition-shadow duration-300">
            <img 
                src="https://media-gru2-1.cdn.whatsapp.net/v/t61.24694-24/420662490_893772325943858_4395920032008760631_n.jpg?ccb=11-4&oh=01_Q5AaIAMmywwrIHcPoiVyrRFb_NHgZHXPK46xiSt_pEXOjaTy&oe=671FE621&_nc_sid=5e03e0&_nc_cat=110" // Substitua pelo caminho da sua imagem
                alt="Desde que te conheci..."
                className="w-full object-contain rounded-t-xl max-h-lvh flex mx-auto"
            />
            <CardContent>
                <h3 className="text-lg font-semibold text-center mb-2">E é claro</h3>
                <p className="text-gray-600 text-center">Com a maior gatinha de todas:</p>

            </CardContent>
            </Card>
        </CarouselItem>

        <CarouselItem className="flex-none w-full px-2">
        <Card className="bg-pink-100 rounded-xl shadow-md hover:shadow-lg transition-shadow duration-300">
            <img 
                src="https://cdn.discordapp.com/attachments/1058805305112285258/1297255317343834112/IMG-20240922-WA0019.jpg?ex=671542a9&is=6713f129&hm=fe85389cefee4800c4f8bffadac5a2f3cb18aede2e79f1191c4f78ceaac265a4&" // Substitua pelo caminho da sua imagem
                alt="Desde que te conheci..."
                className="w-96 object-contain rounded-t-xl max-h-lvh flex mx-auto"
            />
            <CardContent>
                <h3 className="text-lg font-semibold text-center mb-2">Só quero dizer que você me faz muito feliz</h3>
                <p className="text-gray-600 text-center">E que eu eu gostaria muito de continuar assim pro resto da vida</p>
                <p className="text-gray-600 text-center">Mas pra isso eu preciso saber se...</p>


            </CardContent>
            </Card>
        </CarouselItem>

        <CarouselItem className="flex-none w-full px-2">
        <Card className="bg-pink-100 rounded-xl shadow-md hover:shadow-lg transition-shadow duration-300">
            <img 
                src="https://cdn.discordapp.com/attachments/1058805305112285258/1297256463126888549/Cursed_Heart_Images.gif?ex=671543ba&is=6713f23a&hm=1f0fc85b77220001b8fd9ec57a9eaf8cd23719877a94c8c81f751498d3ca1ea0&" // Substitua pelo caminho da sua imagem
                alt="Desde que te conheci..."
                className="w-96 object-contain rounded-t-xl max-h-lvh flex mx-auto"
            />
            <CardContent>
                <h3 className="text-lg font-semibold text-center mb-2">Você quer Namorar comigo?</h3>
                <p className="text-gray-600 text-center">ps: se aceitar vai ter que ser mãe do maluco do osvaldo</p>
                <p className="text-gray-600 text-center">ps2: se não aceitar eu pego a faca de dentro da mochila</p>

               

            </CardContent>
            </Card>
        </CarouselItem>

      </CarouselContent>
      <CarouselNext />
      <CarouselPrevious />
    </Carousel>
  );


  return (
    <div className="p-6">
      {/* Card de Informações do Usuário */}
      {user && (
        <Card className="bg-white rounded-xl shadow-md mb-6">
          <CardContent className="flex items-center">
            <img
              src={user.photoUrl || "/default-user.png"}
              alt={`Foto de ${user.name}`}
              className="w-20 h-20 rounded-full mr-4 object-cover shadow-md"
            />
            <div>
              <CardTitle className="text-2xl font-bold mb-1">{user.name}</CardTitle>
              <CardDescription className="text-gray-600">{user.email}</CardDescription>
              <Button variant="outline" className="mt-3">
                <FontAwesomeIcon icon={faInfoCircle} className="mr-2" /> Ver Perfil Completo
              </Button>
            </div>
          </CardContent>
        </Card>
      )}

      {/* Seção de Pets do Usuário */}
      <div className="mb-6">
        <h2 className="text-xl font-bold mb-4">Meus Pets</h2>
        <Carousel className="w-full max-w-5xl mx-auto">
          <CarouselContent className="flex">
            {pets.map((pet) => (
              <CarouselItem key={pet.id} className="flex-none w-1/3 px-2">
                <Card className="bg-white rounded-xl shadow-md hover:shadow-lg transition-shadow duration-300">
                  <CardContent>
                    {/* Ícone de Espécie */}
                    <div className="flex justify-center mt-2 mb-4">
                      {pet.species === "Cachorro" ? (
                        <FontAwesomeIcon icon={faDog} className="w-6 h-6 text-gray-700" />
                      ) : pet.species === "Gato" ? (
                        <FontAwesomeIcon icon={faCat} className="w-6 h-6 text-gray-700" />
                      ) : null}
                    </div>

                    {/* GIF ou Imagem do Pet */}
                    <img
                      src={pet.image || "/default-pet.png"}
                      alt={`Imagem de ${pet.name}`}
                      className="w-full h-48 object-cover rounded-t-lg"
                    />
                    <div className="p-4">
                      <h3 className="text-lg font-semibold mb-2">{pet.name}</h3>
                      <p className="text-gray-600 mb-1"><strong>Espécie:</strong> {pet.species}</p>
                      <p className="text-gray-600 mb-1"><strong>Raça:</strong> {pet.breed}</p>
                      <p className="text-gray-600 mb-1"><strong>Idade:</strong> {pet.age} anos</p>
                    </div>
                  </CardContent>
                </Card>
              </CarouselItem>
            ))}
          </CarouselContent>
          <CarouselNext />
          <CarouselPrevious />
        </Carousel>
      </div>

      {/* Seção de Consultas */}
      <div className="mb-6">
        <h2 className="text-xl font-bold mb-4">Consultas Recentes</h2>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {consultations.length > 0 ? (
            consultations.map((consultation) => (
              <Card key={consultation.id} className="bg-white rounded-xl shadow-md hover:shadow-lg transition-shadow duration-300">
                <CardContent>
                  <div className="p-4">
                    <h3 className="text-lg font-semibold mb-2">Consulta em {new Date(consultation.date).toLocaleDateString()}</h3>
                    <p className="text-gray-600 mb-1"><strong>Motivo:</strong> {consultation.reason}</p>
                    <p className="text-gray-600 mb-3"><strong>Valor:</strong> R$ {consultation.value.toFixed(2)}</p>
                  </div>
                </CardContent>
              </Card>
            ))
          ) : (
            <p className="text-gray-500">Nenhuma consulta encontrada.</p>
          )}
        </div>
      </div>


    <RomanceCarousel />

    </div>
  );
};

export default UserPage;
