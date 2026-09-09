export default function Home() {
  const tours = [
    { slug: 'fethiye-12-adalar', title: 'Fethiye 12 Adalar Tekne Turu', days: '1 Gün', price: '₺1.700' },
    { slug: 'kapadokya-balon', title: 'Kapadokya Balon Turu', days: '1 Gün', price: '₺12.900' },
    { slug: 'antalya-kultur', title: 'Antalya Kültür Turu', days: '5 Gün', price: '₺15.900' },
  ];
  return (
    <main className="min-h-screen bg-white">
      <header className="p-6 text-center">
        <h1 className="text-4xl font-bold">Keşfet, Rezervasyon Yap</h1>
        <p className="mt-2 text-gray-600">Modern, hızlı ve mobil uyumlu seyahat acentası</p>
      </header>
      <section className="p-6">
        <div className="grid md:grid-cols-3 gap-4">
          {tours.map(t=>(
            <a key={t.slug} href={`/turlar/${t.slug}`} className="border rounded-xl p-4 shadow hover:shadow-md">
              <div className="h-40 bg-gray-200 rounded mb-2"></div>
              <h3 className="font-semibold">{t.title}</h3>
              <p className="text-sm text-gray-500">{t.days}</p>
              <p className="mt-2 font-bold">{t.price}</p>
            </a>
          ))}
        </div>
      </section>
    </main>
  );
}
