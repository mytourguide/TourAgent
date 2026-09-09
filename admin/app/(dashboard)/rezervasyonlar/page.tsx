export default function Rezervasyonlar() {
  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4">Rezervasyon Yönetimi</h1>
      <div className="bg-white rounded-xl shadow p-4">
        <p className="text-gray-500">Filtreler: durum, tur, tarih</p>
        <table className="w-full mt-4">
          <thead><tr className="text-left text-sm"><th>No</th><th>Müşteri</th><th>Tur</th><th>Tutar</th><th>Durum</th></tr></thead>
          <tbody><tr className="border-t"><td>TA-001</td><td>Ahmet Y.</td><td>Kapadokya</td><td>₺25.800</td><td>pending</td></tr></tbody>
        </table>
      </div>
    </div>
  );
}