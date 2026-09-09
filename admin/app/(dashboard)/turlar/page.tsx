export default function Turlar() {
  return (
    <div className="p-6">
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-2xl font-bold">Tur Yönetimi</h1>
        <button className="bg-indigo-600 text-white px-4 py-2 rounded">Yeni Tur</button>
      </div>
      <table className="w-full bg-white rounded-xl shadow overflow-hidden">
        <thead className="bg-gray-100 text-left">
          <tr><th className="p-3">Başlık</th><th>Kategori</th><th>Fiyat</th><th>Durum</th><th></th></tr>
        </thead>
        <tbody>
          <tr className="border-t"><td className="p-3">Kapadokya Balon</td><td>Doğa</td><td>₺12.900</td><td>Aktif</td><td><button className="text-indigo-600">Düzenle</button></td></tr>
        </tbody>
      </table>
    </div>
  );
}