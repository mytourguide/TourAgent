"use client";
import { useState } from "react";
import { useRouter } from "next/navigation";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const router = useRouter();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError("");
    try {
      const base = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api";
      const res = await fetch(`${base}/auth/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });
      if (!res.ok) { const err = await res.text(); throw new Error(err); }
      const data = await res.json();
      localStorage.setItem("token", data.token);
      router.push("/dashboard");
    } catch (e: any) {
      setError(e.message);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="bg-white p-8 rounded-xl shadow w-full max-w-sm">
        <h1 className="text-xl font-bold mb-4">Admin Giriş</h1>
        <form onSubmit={handleSubmit} className="space-y-3">
          <input className="w-full border rounded p-2" placeholder="E-posta" value={email} onChange={e=>setEmail(e.target.value)} />
          <input className="w-full border rounded p-2" type="password" placeholder="Şifre" value={password} onChange={e=>setPassword(e.target.value)} />
          <button className="w-full bg-indigo-600 text-white py-2 rounded" type="submit">Giriş</button>
          {error && <p className="text-red-500 text-sm">{error}</p>}
        </form>
      </div>
    </div>
  );
}
