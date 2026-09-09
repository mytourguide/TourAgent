'use client';
import { useState } from "react";
export default function ImageSlider({ images }:{images:string[]}) {
  const [i,setI]=useState(0);
  return (
    <div className="relative">
      <img src={images[i]} className="w-full h-80 object-cover rounded-xl" alt="" />
      <button onClick={()=>setI((i-1+images.length)%images.length)} className="absolute left-2 top-1/2 -translate-y-1/2 bg-black/40 text-white px-2 rounded">‹</button>
      <button onClick={()=>setI((i+1)%images.length)} className="absolute right-2 top-1/2 -translate-y-1/2 bg-black/40 text-white px-2 rounded">›</button>
    </div>
  );
}
