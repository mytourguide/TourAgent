export default function Sidebar(){
  const items = ["Dashboard","Turlar","Kalkış Tarihleri","Rezervasyonlar","Ödemeler","Müşteriler","Kategoriler","Kuponlar","Menüler"];
  return (
    <aside className="hidden md:block w-64 bg-white border-r h-screen p-4">
      <div className="font-bold text-lg mb-6">Admin</div>
      <nav className="space-y-2">
        {items.map(i=> <a key={i} className="block py-2 px-3 rounded hover:bg-gray-100">{i}</a>)}
      </nav>
    </aside>
  );
}