import "./globals.css";
import Header from "@/components/Header";
import Footer from "@/components/Footer";
export const metadata = { title: "Seyahat Acentası", description: "Modern tur rezervasyon sitesi" };
export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="tr">
      <body>
        <Header />
        {children}
        <Footer />
      </body>
    </html>
  );
}
