const BASE = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080/api";
export async function apiFetch(path:string, opts:any={}) {
  const res = await fetch(`${BASE}${path}`, {
    headers: {"Content-Type":"application/json", ...(opts.headers||{})},
    ...opts
  });
  if(!res.ok) throw new Error(await res.text());
  return res.json();
}