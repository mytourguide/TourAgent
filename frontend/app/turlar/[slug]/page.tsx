import ImageSlider from "@/components/ImageSlider";
import { apiFetch } from "@/lib/api";
import { Metadata } from "next";
export async function generateStaticParams() {
  return [
    { slug: 'fethiye-12-adalar' },
    { slug: 'kapadokya-balon' },
    { slug: 'antalya-kultur' },
  ];
}
export async function generateMetadata({params}:{params:{slug:string}}):Promise<Metadata>{
  return {title:`Tur - ${params.slug}`, description:`${params.slug} tur detayları`};
}
export default async function TurDetay({ params }: { params:{slug:string} }) {
  let tour:any=null;
  try { tour = await apiFetch(`/tours/${params.slug}`); } catch {}
  const images = tour?.images || [
    "https://images.unsplash.com/photo-1507525428034-b723cf961d3e?w=1200",
    "https://images.unsplash.com/photo-1510414842594-a61c69b5ae57?w=1200",
    "https://images.unsplash.com/photo-1501554728187-ce583db33af7?w=1200",
    "https://images.unsplash.com/photo-1491555103944-7c647fd857e6?w=1200",
  ];
  const title = tour?.title || "Fethiye 12 Adalar Tekne Turu";
  return (
    <main className="p-6 max-w-6xl mx-auto">
      <h1 className="text-3xl font-bold">{title}</h1>
      <ImageSlider images={images} />
      <div className="grid md:grid-cols-3 gap-6 mt-6">
        <div className="md:col-span-2 space-y-6">
          <section>
            <h2 className="text-xl font-semibold mb-2">Tur Açıklaması</h2>
            <p className="text-gray-600">Ege ve Akdeniz'in incisi Fethiye'den kalkan 12 Adalar tekne turu. Öğle yemeği, yüzme molaları ve adalar arasında keşif.</p>
          </section>
          <section>
            <h3 className="font-semibold mb-2">Tura Dahil Olanlar</h3>
            <ul className="list-disc pl-5 text-green-700">{["Tekne ulaşımı","Öğle yemeği","Rehberlik","Sigorta"].map(i=><li key={i}>{i}</li>)}</ul>
          </section>
          <section>
            <h3 className="font-semibold mb-2">Tura Dahil Olmayanlar</h3>
            <ul className="list-disc pl-5 text-red-700">{["Kişisel harcamalar","İçecekler","Fotoğraf ücreti"].map(i=><li key={i}>{i}</li>)}</ul>
          </section>
          <section>
            <h3 className="font-semibold mb-2">Yanınızda Getirin</h3>
            <ul className="list-disc pl-5">{["Nüfus cüzdanı","Mayo","Şapka","Güneş kremi"].map(i=><li key={i}>{i}</li>)}</ul>
          </section>
        </div>
        <aside className="border rounded-xl p-4 h-fit sticky top-6">
          <p className="text-sm text-gray-500">Tarih seç</p>
          <select className="w-full border rounded p-2 mt-2"><option>12 Oct 2026</option></select>
          <div className="mt-4 space-y-3">
            <div>
              <label className="text-sm">Yetişkin</label>
              <div className="flex items-center gap-2"><button className="border px-2">-</button><span>2</span><button className="border px-2">+</button></div>
            </div>
            <div>
              <label className="text-sm">Çocuk</label>
              <div className="flex items-center gap-2"><button className="border px-2">-</button><span>0</span><button className="border px-2">+</button></div>
            </div>
            <p className="font-bold">Toplam: ₺1.700</p>
          </div>
          <button className="mt-4 w-full bg-indigo-600 text-white py-2 rounded">Sepete Ekle</button>
          <div className="mt-4 text-xs text-gray-500">Kişisel bilgiler ödeme adımında istenecektir.</div>
        </aside>
      </div>
    </main>
  );
}
