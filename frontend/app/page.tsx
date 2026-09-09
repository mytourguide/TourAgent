export default function Home() {
  const tours = [
    { slug: 'fethiye-12-adalar', title: 'Fethiye 12 Adalar Tekne Turu', days: '1 Gün', price: '₺1.700', img: 'https://images.unsplash.com/photo-1507525428034-b723cf961d3e?w=600' },
    { slug: 'kapadokya-balon', title: 'Kapadokya Balon Turu', days: '1 Gün', price: '₺12.900', img: 'https://images.unsplash.com/photo-1541005735092-9b5a7a5b3e5a?w=600' },
    { slug: 'antalya-kultur', title: 'Antalya Kültür Turu', days: '5 Gün', price: '₺15.900', img: 'https://images.unsplash.com/photo-1501554728187-ce583db33af7?w=600' },
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
            <a key={t.slug} href={`/turlar/${t.slug}`} className="block border rounded-xl p-4 shadow hover:shadow-md">
              <img src={t.img} alt={t.title} className="w-full h-40 object-cover rounded mb-2" />
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
