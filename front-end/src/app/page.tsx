import type { Metadata } from "next";

export default function Home() {
  return (
    <div className="pt-16 min-h-screen bg-gray-100 flex items-center justify-center">
      <div className="text-center">
        <h1 className="text-4xl font-bold text-gray-800 mb-4">
          Bem-vindo ao Figueiredo Imóveis
        </h1>
        <p className="text-lg text-gray-600">
          Sistema de gestão imobiliária com inteligência artificial.
        </p>
        <p className="text-sm text-gray-500 mt-2">
          Está funcionando! 🚀
        </p>
      </div>
    </div>
  );
}
