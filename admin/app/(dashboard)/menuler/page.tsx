export default function Menuler(){
  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4">Site Menü Yönetimi</h1>
      <div className="bg-white rounded-xl shadow p-4">
        <p className="text-sm text-gray-500 mb-4">Menü tipleri, sıralama, aktif/pasif ve alt menü düzenlemeleri burada yapılır.</p>
        <table className="w-full text-sm">
          <thead><tr className="text-left border-b"><th className="p-2">Ad</th><th>URL</th><th>Sıra</th><th>Aktif</th></tr></thead>
          <tbody><tr className="border-b"><td className="p-2">Ana Sayfa</td><td>/</td><td>1</td><td>✓</td></tr></tbody>
        </table>
        <button className="mt-4 bg-indigo-600 text-white px-4 py-2 rounded">Yeni Menü</button>
      </div>
    </div>
  );
}