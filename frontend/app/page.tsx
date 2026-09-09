export default function Home() {
  return (
    <main className="min-h-screen bg-white">
      <header className="p-6 text-center">
        <h1 className="text-4xl font-bold">Keşfet, Rezervasyon Yap</h1>
        <p className="mt-2 text-gray-600">Modern, hızlı ve mobil uyumlu seyahat acentası</p>
      </header>
      <section className="p-6">
        <div className="grid md:grid-cols-3 gap-4">
          {[1,2,3].map(i=>(
            <div key={i} className="border rounded-xl p-4 shadow">
              <div className="h-40 bg-gray-200 rounded mb-2"></div>
              <h3 className="font-semibold">Öne Çıkan Tur {i}</h3>
              <p className="text-sm text-gray-500">7 Gün 6 Gece · İstanbul</p>
              <p className="mt-2 font-bold">₺12.900</p>
            </div>
          ))}
        </div>
      </section>
    </main>
  );
}