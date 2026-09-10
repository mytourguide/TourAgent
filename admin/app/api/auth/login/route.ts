import { NextRequest, NextResponse } from "next/server";

export async function POST(req: NextRequest) {
  try {
    const { email, password } = await req.json();
    const base = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api";
    const res = await fetch(`${base}/auth/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    });
    if (!res.ok) {
      const err = await res.text();
      return NextResponse.json({ error: err }, { status: res.status });
    }
    const data = await res.json();
    const resp = NextResponse.json({ success: true, user: data.user });
    resp.cookies.set("token", data.token, { httpOnly: true, secure: false, sameSite: "lax", maxAge: 60 * 60 * 24 });
    return resp;
  } catch (e: any) {
    return NextResponse.json({ error: e.message }, { status: 500 });
  }
}
