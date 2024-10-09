import { useState, useEffect } from "react";
import Link from "next/link";
import { Bell, Home, ShoppingCart, LineChart, Users } from "lucide-react";
import Image from "next/image";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faDog, faNotesMedical, faUser } from "@fortawesome/free-solid-svg-icons"; // Adicionando o ícone de perfil
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuLabel, DropdownMenuSeparator, DropdownMenuTrigger } from "@/components/ui/dropdown-menu";

export const Navbar = () => {
  const [isOpen, setIsOpen] = useState(false);
  const [isMobile, setIsMobile] = useState(false);

  const toggleNavbar = () => setIsOpen((prev) => !prev);

  useEffect(() => {
    const handleResize = () => setIsMobile(window.innerWidth < 768);
    handleResize();
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

  return isMobile ? (
    <div>
      <Button variant="outline" size="icon" onClick={toggleNavbar}>
        <FontAwesomeIcon icon={faDog} />
      </Button>
      {isOpen && (
        <div className="fixed top-0 left-0 w-72 h-full bg-slate-500 text-white">
          <nav className="p-4 gap-4">
            <Button variant="outline" size="icon" onClick={toggleNavbar} className="h-10 w-10 hover:scale-125 flex mb-3">
              <FontAwesomeIcon icon={faDog} className="text-black" />
            </Button>
            <Link href="#" className="flex items-center">
              <Home className="h-5 w-5 mr-2" /> Dashboard
            </Link>
            <Link href="#" className="flex items-center">
              <ShoppingCart className="h-5 w-5 mr-2" /> Orders
            </Link>
            <Link href="#" className="flex items-center">
              <FontAwesomeIcon icon={faNotesMedical} className="h-5 w-5 mr-2" /> Products
            </Link>
            <Link href="#" className="flex items-center">
              <Users className="h-5 w-5 mr-2" /> Customers
            </Link>
            <Link href="#" className="flex items-center">
              <LineChart className="h-5 w-5 mr-2" /> Analytics
            </Link>
          </nav>

          {/* Ícone de perfil fixado no canto inferior esquerdo */}
          <div className="absolute bottom-0 left-0 w-full p-4 flex items-center">
            <FontAwesomeIcon icon={faUser} className="h-6 w-6 text-white mr-2" />
            {isOpen && <span className="text-white">Perfil</span>}
          </div>
        </div>
      )}
    </div>
  ) : (
    <div className={`fixed left-0 top-0 h-full border-r bg-muted transition-all duration-300 ${isOpen ? 'w-[220px]' : 'w-[80px]'}`}>
      <div className="flex h-full flex-col gap-2">
        <div className="flex h-14 items-center border-b px-4 lg:h-[60px] lg:px-6">
          <Button variant="outline" size="icon" onClick={toggleNavbar} className="h-10 w-10 hover:scale-125">
            <FontAwesomeIcon icon={faDog} />
          </Button>
          <Link href="/" className="flex items-center gap-2 font-semibold">
            {isOpen && <span className="ml-2">Vetsys</span>}
          </Link>
        </div>
        <nav className="flex-1 px-2 text-sm font-medium lg:px-4">
          <Link href="#" className="flex items-center gap-3 px-3 py-2 hover:bg-muted transition-colors rounded-lg">
            <Home className="h-4 w-4" />
            {isOpen && <span>Dashboard</span>}
          </Link>
          <Link href="#" className="flex items-center gap-3 px-3 py-2 hover:bg-muted transition-colors rounded-lg">
            <ShoppingCart className="h-4 w-4" />
            {isOpen && <span>Orders</span>}
            <Badge className={`ml-auto h-6 w-6 ${isOpen ? '' : 'hidden'}`}>6</Badge>
          </Link>
          <Link href="/create-consultation" className="flex items-center gap-3 px-3 py-2 hover:bg-muted transition-colors rounded-lg">
            <FontAwesomeIcon icon={faNotesMedical} className="h-4 w-4" />
            {isOpen && <span>Products</span>}
          </Link>
          <Link href="#" className="flex items-center gap-3 px-3 py-2 hover:bg-muted transition-colors rounded-lg">
            <Users className="h-4 w-4" />
            {isOpen && <span>Customers</span>}
          </Link>
          <Link href="#" className="flex items-center gap-3 px-3 py-2 hover:bg-muted transition-colors rounded-lg">
            <LineChart className="h-4 w-4" />
            {isOpen && <span>Analytics</span>}
          </Link>
        </nav>

        {/* Ícone de perfil fixado no canto inferior esquerdo */}
        <div className="relative mt-auto p-4">
          {/* bordas redondas icone */}
          {/* Ícone de perfil com bordas redondas */}
          <div className="absolute bottom-0 left-0 w-full flex items-center justify-center p-4">
          {/* Ícone com fundo, bordas redondas e alinhado */}
          <div className="h-12 w-12 bg-gray-200 rounded-full flex items-center justify-center">
          <DropdownMenu>
            <DropdownMenuTrigger asChild>
              <Button
                variant="outline"
                size="icon"
                className="overflow-hidden rounded-full"
              >
               <img src="https://avatar.iran.liara.run/public/boy?username=Ash" alt="" />
              </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              <DropdownMenuLabel>My Account</DropdownMenuLabel>
              <DropdownMenuSeparator />
              <DropdownMenuItem>Settings</DropdownMenuItem>
              <DropdownMenuItem>Support</DropdownMenuItem>
              <DropdownMenuSeparator />
              <DropdownMenuItem>Logout</DropdownMenuItem>
            </DropdownMenuContent>
          </DropdownMenu>
          </div>
          {isOpen && <span className="ml-2">Perfil</span>}
        </div>
        </div>
      </div>
    </div>
  );
};
