"use client";
import { useEffect, useState } from "react";

interface Stat { label: string; value: string; }

export default function Dashboard() {
  const [stats, setStats] = useState<Stat[]>([]);
  const base = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api";
  useEffect(() => {
    fetch(`${base}/tours`).then(r=>r.json()).then(data=>{
      setStats([
        { label:"Toplam Rezervasyon", value:"—" },
        { label:"Bekleyen Ödeme", value:"—" },
        { label:"Aylık Ciro", value:"—" },
        { label:"Toplam Tur", value:String(data.data?.length||0) },
      ]);
    }).catch(()=>{});
  }, [base]);
  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4">Dashboard</h1>
      <div className="grid md:grid-cols-4 gap-4">
        {stats.map(s=>(
          <div key={s.label} className="bg-white rounded-xl shadow p-4">
            <p className="text-sm text-gray-500">{s.label}</p>
            <p className="text-2xl font-bold mt-2">{s.value}</p>
          </div>
        ))}
      </div>
    </div>
  );
}
