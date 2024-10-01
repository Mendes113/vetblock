import { useState, useEffect } from "react";
import Link from "next/link";
import { Bell, Home, ShoppingCart, LineChart, Users, CircleUser } from "lucide-react";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faBars, faDog, faNotesMedical, faStar, faTimes } from "@fortawesome/free-solid-svg-icons";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator } from "../ui/dropdown-menu";
import { DropdownMenuTrigger } from "@radix-ui/react-dropdown-menu";

export const Navbar = () => {
  const [isOpen, setIsOpen] = useState(false); // Começa fechado em telas menores
  const [isRotating, setIsRotating] = useState(false);

  const toggleNavbar = () => {
    setIsOpen((prev) => !prev); // Alterna o estado de isOpen
  };

  const handleUpgradeClick = () => {
    setIsRotating(true);
    setTimeout(() => {
      setIsRotating(false);
    }, 1000); 
  };

  useEffect(() => {
    // Função para atualizar o estado da navbar baseado no tamanho da tela
    const handleResize = () => {
      if (window.innerWidth < 990) {
        setIsOpen(false); // Fecha a navbar em telas menores
      } 
    };

    // Executa a função uma vez quando o componente monta
    handleResize();

    // Adiciona um listener para o redimensionamento da tela
    window.addEventListener("resize", handleResize);

    // Remove o listener quando o componente desmonta
    return () => window.removeEventListener("resize", handleResize);
  }, []);


  const isMobile = window.innerWidth < 768;
  return (
    !isMobile ? (

    <div className={`fixed left-0 top-0 h-full border-r bg-muted/40 transition-all duration-300 ease-in-out ${isOpen ? 'w-[220px]' : 'w-[80px] bg-slate-100'}`}>
      <div className="flex h-full max-h-screen flex-col gap-2">
        <div className="flex h-14 items-center border-b px-4 lg:h-[60px] lg:px-6">
          <Button variant="outline" size="icon"
            className="border-none h-10 w-10 transition-transform duration-300 ease-in-out transform hover:scale-125"
            onClick={toggleNavbar}>
            <FontAwesomeIcon icon={faDog} />
          </Button>
          <Link href="/" className="flex items-center gap-2 font-semibold">
            {isOpen && <span className="ml-2">Vetsys</span>}
          </Link>
          <Button variant="outline" size="icon" className={`ml-auto h-8 w-8 ${isOpen ? '' : 'md:hidden'}`}>
            <Bell className="h-4 w-4" />
            <span className="sr-only">Toggle notifications</span>
          </Button>
        </div>
        <div className="flex-1">
          <nav className="grid items-start px-2 text-sm font-medium lg:px-4">
            <Link
              href="#"
              className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
            >
              <Home className="h-4 w-4" />
              {isOpen && <span>Dashboard</span>}
            </Link>
            <Link
              href="#"
              className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
            >
              <ShoppingCart className="h-4 w-4" />
              {isOpen && <span>Orders</span>}
              <Badge className={`ml-auto flex h-6 w-6 shrink-0 items-center justify-center rounded-full ${isOpen ? '' : 'hidden'}`}>
                6
              </Badge>
            </Link>
            <Link
              href="#"
              className="flex items-center gap-3 rounded-lg bg-muted px-3 py-2 text-primary transition-all hover:text-primary"
            >
              <FontAwesomeIcon icon={faNotesMedical} className="h-4 w-4" />
              {isOpen && <span>Products</span>}
            </Link>
            <Link
              href="#"
              className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
            >
              <Users className="h-4 w-4" />
              {isOpen && <span>Customers</span>}
            </Link>
            <Link
              href="#"
              className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:text-primary"
            >
              <LineChart className="h-4 w-4" />
              {isOpen && <span>Analytics</span>}
            </Link>
          </nav>
        </div>
        <div className="mt-auto p-4">
          {isOpen ? (
            <div>
              <Card>
                <CardHeader className="p-2 pt-0 md:p-4">
                  <CardTitle>Upgrade to Pro</CardTitle>
                  <CardDescription>
                    Unlock all features and get unlimited access to our support team.
                  </CardDescription>
                </CardHeader>
                <CardContent className="p-2 pt-0 md:p-4 md:pt-0">
                  <Button size="sm" className="w-full" onClick={handleUpgradeClick}>
                    Upgrade
                  </Button>
                </CardContent>
              </Card>
              <Button variant="outline" size="icon" className="ml-auto h-8 w-8 mt-2">
                <Bell className="h-4 w-4" />
                <span className="sr-only">Toggle notifications</span>
              </Button>
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button variant="outline" className="h-14 w-14 rounded-full mt-2">
                    <CircleUser className="h-4 w-4" />
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                  <DropdownMenuLabel>User settings</DropdownMenuLabel>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem>Profile</DropdownMenuItem>
                  <DropdownMenuItem>Billing</DropdownMenuItem>
                  <DropdownMenuItem>Settings</DropdownMenuItem>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem>Log out</DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            </div>
          ) : (
            <div className="flex-col justify-between">
              <Button variant="outline" size="icon" className="w-full" onClick={toggleNavbar}>
                <FontAwesomeIcon icon={faStar} />
              </Button>
              <Button variant="outline" size="icon" className="ml-auto h-8 w-8 mt-2">
                <Bell className="h-4 w-4" />
                <span className="sr-only">Toggle notifications</span>
              </Button>
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button variant="outline" className="h-12 w-12 rounded-full mt-2">
                    <CircleUser className="h-8 w-8" />
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                  <DropdownMenuLabel>User settings</DropdownMenuLabel>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem>Profile</DropdownMenuItem>
                  <DropdownMenuItem>Billing</DropdownMenuItem>
                  <DropdownMenuItem>Settings</DropdownMenuItem>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem>Log out</DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            </div>
          )}
        </div>
      </div>
    </div>
  ) : (
        <div>
          

          <div>
    <Button variant="outline" size="icon" onClick={toggleNavbar}>
      <FontAwesomeIcon icon={isOpen ? faDog : faDog} />
    </Button>

    {isOpen && (
      <div className="fixed flex-col top-0 left-0 w-72 h-full bg-slate-500 text-white">
        <nav className="flex flex-col  justify-self-start gap-4 p-4 w-52 ">
          <Link href="#" className="text-lg flex ">
            <Home className="h-5 w-5 mr-2" /> Dashboard
          </Link>
          <Link href="#" className="text-lg flex">
            <ShoppingCart className="h-5 w-5 mr-2" /> Orders
          </Link>
          <Link href="#" className="text-lg flex ">
            <FontAwesomeIcon icon={faNotesMedical} className="h-5 w-5 mr-2" /> Products
          </Link>
          <Link href="#" className="text-lg flex">
            <Users className="h-5 w-5 mr-2" /> Customers
          </Link>
          <Link href="#" className="text-lg flex">
            <LineChart className="h-5 w-5 mr-2" /> Analytics
          </Link>
        </nav>

        {/* pro */}
        <div className="flex-col justify-between">
              <Button variant="outline" size="icon" className="w-full" onClick={toggleNavbar}>
                <FontAwesomeIcon icon={faStar} />
              </Button>
              <Button variant="outline" size="icon" className="ml-auto h-8 w-8 mt-2">
                <Bell className="h-4 w-4" />
                <span className="sr-only">Toggle notifications</span>
              </Button>
              <DropdownMenu>
                <DropdownMenuTrigger asChild>
                  <Button variant="outline" className="h-12 w-12 rounded-full mt-2">
                    <CircleUser className="h-8 w-8" />
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent align="end">
                  <DropdownMenuLabel>User settings</DropdownMenuLabel>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem>Profile</DropdownMenuItem>
                  <DropdownMenuItem>Billing</DropdownMenuItem>
                  <DropdownMenuItem>Settings</DropdownMenuItem>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem>Log out</DropdownMenuItem>
                </DropdownMenuContent>
              </DropdownMenu>
            </div>

      </div>
    )}
  </div>

 
        
        </div>
    )
    );
};
