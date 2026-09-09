export default function Dashboard() {
  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4">Dashboard</h1>
      <div className="grid md:grid-cols-4 gap-4">
        {["Toplam Rezervasyon","Bekleyen Ödeme","Aylık Ciro","Doluluk Oranı"].map(k=>(
          <div key={k} className="bg-white rounded-xl shadow p-4">
            <p className="text-sm text-gray-500">{k}</p>
            <p className="text-2xl font-bold mt-2">—</p>
          </div>
        ))}
      </div>
      <div className="mt-6 bg-white rounded-xl shadow p-4 h-64">Grafik alanı</div>
    </div>
  );
}