import TourCard from "@/components/TourCard";
import { apiFetch } from "@/lib/api";
import type { Tour } from "@/lib/types";

async function getTours(): Promise<Tour[]> {
  try {
    const data = await apiFetch("/tours?limit=12");
    return data.data || [];
  } catch { return []; }
}

export default async function Turlar() {
  const tours = await getTours();
  const dummy = tours.length ? tours : [
    {id:"1",title:"Kapadokya Balon Turu", slug:"kapadokya-balon", location:"Nevşehir", duration_days:3, base_price:12900, cover_image:"https://picsum.photos/seed/1/600/400"},
    {id:"2",title:"Antalya Kültür Turu", slug:"antalya-kultur", location:"Antalya", duration_days:5, base_price:15900, cover_image:"https://picsum.photos/seed/2/600/400"},
  ];
  return (
    <main className="p-6 max-w-7xl mx-auto">
      <h1 className="text-3xl font-bold mb-6">Turlar</h1>
      <div className="flex flex-wrap gap-2 mb-6">
        {["Tümü","Mavi Turlar","Extrem Turlar","Aile Turları","Macera Turları","Kültür","Doğa"].map(c=>(
          <button key={c} className="px-3 py-1 border rounded-full text-sm hover:bg-indigo-50">{c}</button>
        ))}
      </div>
      <div className="grid sm:grid-cols-2 lg:grid-cols-3 gap-6">
        {dummy.map(t=> <TourCard key={t.slug} title={t.title} slug={t.slug} location={t.location} days={t.duration_days||t.days||3} price={t.base_price||t.price||0} image={t.cover_image||t.image||""} />)}
      </div>
    </main>
  );
}