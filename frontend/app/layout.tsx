import "./globals.css";
import Header from "@/components/Header";
import Footer from "@/components/Footer";
export const metadata = { title: "Seyahat Acentası", description: "Modern tur rezervasyon sitesi" };
export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="tr">
      <head>
        <script src="https://cdn.tailwindcss.com"></script>
      </head>
      <body>
        <Header />
        {children}
        <Footer />
      </body>
    </html>
  );
}
