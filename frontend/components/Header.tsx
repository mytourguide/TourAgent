'use client';
import { useEffect, useState } from "react";
export default function Header(){
  const [menus,setMenus]=useState<{name:string;url:string}[]>([]);
  useEffect(()=>{fetch(process.env.NEXT_PUBLIC_API_URL+"/api/admin/menus").then(r=>r.json()).then(d=>setMenus(d.data||[])).catch(()=>{})},[]);
  return (
    <header className="border-b">
      <div className="max-w-6xl mx-auto p-4 flex justify-between items-center">
        <a href="/" className="font-bold text-xl">Seyahat Acentası</a>
        <nav className="hidden md:flex space-x-4">
          {menus.length?menus.map(m=><a key={m.url} href={m.url}>{m.name}</a>):<>
            <a href="/turlar">Turlar</a>
            <a href="/hesabim">Hesabım</a>
          </>}
        </nav>
      </div>
    </header>
  );
}