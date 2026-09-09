export default function Rezervasyon() {
  return (
    <main className="p-6 max-w-3xl mx-auto">
      <h1 className="text-2xl font-bold mb-4">Rezervasyon</h1>
      <div className="space-y-4">
        <section className="border p-4 rounded">
          <h2 className="font-semibold mb-2">Yolcu Bilgileri</h2>
          <div className="grid sm:grid-cols-2 gap-3">
            <input className="border rounded p-2" placeholder="Ad Soyad" />
            <input className="border rounded p-2" type="date" placeholder="Doğum Tarihi" />
            <input className="border rounded p-2" placeholder="Kimlik/Pasaport No" />
            <select className="border rounded p-2"><option>Yetişkin</option><option>Çocuk</option></select>
          </div>
        </section>
        <section className="border p-4 rounded">
          <h2 className="font-semibold">Özet</h2>
          <p>Tur: Fethiye 12 Adalar · Tarih: 12 Oct 2026 · Yetişkin: 2 · Çocuk: 1 · Toplam: ₺2.550</p>
        </section>
        <section className="border p-4 rounded">
          <h2 className="font-semibold">Ödeme</h2>
          <p className="text-sm text-gray-500">İyzico Checkout Form burada embed edilecek (3D Secure).</p>
          <button className="mt-3 bg-indigo-600 text-white px-4 py-2 rounded">Ödemeyi Başlat</button>
        </section>
      </div>
    </main>
  );
}