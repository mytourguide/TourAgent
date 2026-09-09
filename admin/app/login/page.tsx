export default function LoginPage(){
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-xl shadow w-full max-w-sm">
        <h1 className="text-xl font-bold mb-4">Admin Giriş</h1>
        <form className="space-y-3">
          <input className="w-full border rounded p-2" placeholder="E-posta" />
          <input className="w-full border rounded p-2" type="password" placeholder="Şifre" />
          <button className="w-full bg-indigo-600 text-white py-2 rounded">Giriş</button>
        </form>
      </div>
    </div>
  );
}