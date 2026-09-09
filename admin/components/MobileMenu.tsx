export default function MobileMenu(){
  return (
    <div className="md:hidden fixed top-0 left-0 right-0 bg-white border-b p-3 flex justify-between">
      <span className="font-bold">Admin</span>
      <button className="border px-3 py-1 rounded">☰</button>
    </div>
  );
}