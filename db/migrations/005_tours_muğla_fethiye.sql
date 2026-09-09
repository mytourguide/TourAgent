-- Muğla / Fethiye örnek turları
INSERT INTO tours (id, title, slug, description, category_id, cover_image, duration_days, duration_nights, location, base_price, currency, is_active, is_featured)
SELECT gen_random_uuid(), 'Fethiye 12 Adalar Tekne Turu', 'fethiye-12-adalar-tekne-turu', 'Ege ve Akdeniz''in incisi Fethiye''den kalkan 12 Adalar tekne turu. Öğle yemeği, yüzme molaları ve adalar arasında keşif.', (SELECT id FROM tour_categories WHERE slug='mavi-turlar'), 'https://images.unsplash.com/photo-1507525428034-b723cf961d3e?w=1200', 1,0,'Muğla/Fethiye', 850, 'TRY', true, true
UNION ALL SELECT gen_random_uuid(), 'Ölüdeniz Yamaç Paraşütü', 'oludeniz-yamac-parasutu', 'Dünya''nın en güzel yamaç paraşütü noktalarından Ölüdeniz''de tandem uçuş.', (SELECT id FROM tour_categories WHERE slug='macera-turlar'), 'https://images.unsplash.com/photo-1469474968028-56623f02e42e?w=1200', 1,0,'Muğla/Ölüdeniz', 1800,'TRY',true,true
UNION ALL SELECT gen_random_uuid(), 'Saklıkent Kanyonu Rafting', 'saklikent-kanyonu-rafting', 'Saklıkent Kanyonu''nda adrenalin dolu rafting turu.', (SELECT id FROM tour_categories WHERE slug='extrem-turlar'), 'https://images.unsplash.com/photo-1501785888041-af3ef285b470?w=1200',1,0,'Muğla/Kaş',1200,'TRY',true,false
UNION ALL SELECT gen_random_uuid(), 'Kaunos Köprülü Kanyon Kano', 'kaunos-koopru-ulu-kanyon-kano', 'Likya uygarlığı kalıntıları ve kanyon kano.', (SELECT id FROM tour_categories WHERE slug='aile-turlar'), 'https://images.unsplash.com/photo-1501785888041-af3ef285b470?w=1200',1,0,'Muğla/Dalaman',950,'TRY',true,false;

-- Görseller
INSERT INTO tour_images (tour_id, image_url, sort_order)
SELECT t.id, 'https://images.unsplash.com/photo-1507525428034-b723cf961d3e?w=1200',1 FROM tours t WHERE t.slug='fethiye-12-adalar-tekne-turu'
UNION ALL SELECT t.id, 'https://images.unsplash.com/photo-1510414842594-a61c69b5ae57?w=1200',2 FROM tours t WHERE t.slug='fethiye-12-adalar-tekne-turu'
UNION ALL SELECT t.id, 'https://images.unsplash.com/photo-1501554728187-ce583db33af7?w=1200',3 FROM tours t WHERE t.slug='fethiye-12-adalar-tekne-turu'
UNION ALL SELECT t.id, 'https://images.unsplash.com/photo-1491555103944-7c647fd857e6?w=1200',4 FROM tours t WHERE t.slug='fethiye-12-adalar-tekne-turu';
