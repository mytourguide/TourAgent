"use client";
import { useEffect } from "react";
import { useRouter } from "next/navigation";

export default function DashboardLayout({ children }: { children: React.ReactNode }) {
  const router = useRouter();
  const token = typeof window !== "undefined" ? localStorage.getItem("token") : null;
  useEffect(() => {
    if (!token) router.push("/login");
  }, [token, router]);
  if (!token) return <div className="min-h-screen flex items-center justify-center bg-gray-100">Yükleniyor...</div>;
  return <div className="flex">{children}</div>;
}
