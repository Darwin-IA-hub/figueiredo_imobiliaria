import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import Navbar from "../components/Navbar";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Figueiredo Imóveis",
  description: "Sistema de gestão imobiliária Figueiredo Imóveis",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="pt-BR" className="h-full">
      <body
        className={`min-h-full ${geistSans.variable} ${geistMono.variable} antialiased bg-white dark:bg-[#0a0a0a] text-gray-900 dark:text-gray-100`}
      >
        <Navbar />
        {children}
      </body>
    </html>
  );
}
