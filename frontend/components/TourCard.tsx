type TourCardProps = { title:string; slug:string; location:string; days:number; price:number; image:string };
export default function TourCard({title, slug, location, days, price, image}: TourCardProps) {
  return (
    <a href={`/turlar/${slug}`} className="border rounded-xl overflow-hidden shadow hover:shadow-md transition">
      <div className="h-48 bg-gray-200" style={{backgroundImage:`url(${image})`,backgroundSize:'cover'}}/>
      <div className="p-4">
        <h3 className="font-semibold">{title}</h3>
        <p className="text-sm text-gray-500">{location} · {days} gün</p>
        <p className="mt-2 font-bold">₺{price.toLocaleString('tr-TR')}</p>
      </div>
    </a>
  );
}