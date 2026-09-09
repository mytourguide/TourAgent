import "./globals.css";
export const metadata = { title: "Admin Panel", description: "Seyahat acentası yönetim paneli" };
export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="tr">
      <body className="bg-gray-50">{children}</body>
    </html>
  );
}