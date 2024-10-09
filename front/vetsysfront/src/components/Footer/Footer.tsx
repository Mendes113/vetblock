import { useEffect, useState, useRef } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFacebook, faTwitter, faInstagram } from '@fortawesome/free-brands-svg-icons';
import Link from 'next/link';

export const Footer = () => {
  const [isFooterVisible, setIsFooterVisible] = useState(true);
  const footerRef = useRef(null);
  const navRef = useRef(null);

  useEffect(() => {
    const footer = footerRef.current;
    const nav = navRef.current;

    // Função de callback quando ocorrer uma intersecção
    const handleIntersection = (entries) => {
      entries.forEach((entry) => {
        if (entry.isIntersecting) {
          setIsFooterVisible(false); // Esconde o footer se ele intersectar com a navbar
        } else {
          setIsFooterVisible(true); // Mostra o footer quando não há intersecção
        }
      });
    };

    // Configurando o Intersection Observer
    const observer = new IntersectionObserver(handleIntersection, {
      root: null,
      threshold: 0.1,
    });

    if (footer && nav) {
      observer.observe(nav); // Observa a navbar
      observer.observe(footer); // Observa o footer
    }

    // Cleanup do observer quando o componente desmonta
    return () => {
      if (footer) observer.unobserve(footer);
      if (nav) observer.unobserve(nav);
    };
  }, []);

  return (
    <>
   

      {/* Footer */}
      <footer
        ref={footerRef}
        className={`w-full bg-muted/40 p-4 border-t transition-opacity duration-500 ${
          isFooterVisible ? 'opacity-100' : 'opacity-0'
        }`}
      >
        <div className="container mx-auto flex flex-col md:flex-row justify-between items-center">
          {/* Seção de links */}
          <div className="mb-4 md:mb-0">
            <nav className="flex gap-4">
              <Link href="/" className="text-gray-600 hover:text-gray-900">Home</Link>
              <Link href="/about" className="text-gray-600 hover:text-gray-900">About</Link>
              <Link href="/contact" className="text-gray-600 hover:text-gray-900">Contact</Link>
              <Link href="/privacy" className="text-gray-600 hover:text-gray-900">Privacy Policy</Link>
            </nav>
          </div>

          {/* Seção de mídia social */}
          <div className="flex gap-4">
            <a href="https://facebook.com" target="_blank" rel="noopener noreferrer">
              <FontAwesomeIcon icon={faFacebook} className="text-gray-600 hover:text-blue-600 h-5 w-5" />
            </a>
            <a href="https://twitter.com" target="_blank" rel="noopener noreferrer">
              <FontAwesomeIcon icon={faTwitter} className="text-gray-600 hover:text-blue-400 h-5 w-5" />
            </a>
            <a href="https://instagram.com" target="_blank" rel="noopener noreferrer">
              <FontAwesomeIcon icon={faInstagram} className="text-gray-600 hover:text-pink-500 h-5 w-5" />
            </a>
          </div>

          {/* Seção de copyright */}
          <div className="mt-4 md:mt-0 text-gray-600">
            <span>&copy; {new Date().getFullYear()} Vetsys. Todos os direitos reservados.</span>
          </div>
        </div>
      </footer>
    </>
  );
};
