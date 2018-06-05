# Pokedex
### by Attila Satan

## Giriş

Bu proje boyunca her ne kadar bahsedilemiş olsa da kendime iki hedef daha belirledim. Birincisi **kolay geliştirme** diğeri ise **işlem gücü ihtiyacını azaltma**. Ayrıca olabildiğince Domain Driven Development (DDD) metodolojisine bağlı kalmaya çalıştım.

Çok geniş kapsamlı bir API yazmak yerine nasıl yazılabileceiğini göstermeye çalıştım. Bu aşamadan sonra geliştirme kolay olacaktır.

Kodu okurken mutlu olacağınızı ümit ederim. 

Aradığımız veriyi sliceları gezerek toplayabilirdik çoğu yerde de böyle yaptım ama bazı `string` manipulasyonu gerektiren filtrelemelerde indeksleme metodunu kullandım.
 
Indeksleme yaparak bir kaç kilobyte rem ile microsunucudan tasarruf sağlayabiliriz. 

Geliştirmeyi kolaylaştırmak için `Pokedex` pakelti oluşturup `main` metodunu `cnd` klasörüne taşıdım.

Bu sayede `pokedex` bir API'a sahip olmuş oldu.

Bir çok yerde gereksiz yere public kullandığımın farkındayım. go yazdığımı çok sonra hatırladım. 

End-Point sayısı az görünebilir ancak hazırladığım basit mimari ve `api` sayesinde yeni filtre ve sıralama ölçütü eklemek çok kolay olacaktır.

Mutlu rewievlar.

## End-Points

Filtrelemede genellikle keyler case insensitive olarak tasarlandı ancak %100 garanti vermemekteyim :) 

    /api/pokemon/Charmeleon == /api/pokemon/CHARMELEON == /api/pokemon/charmeleon

Filtre ve sorting sistemi sadece Pokemon resource'u için geliştirildi. Başında `-` bulunan sorting valuelar tersten sıralama yapar.

### /api/pokemon/list
Pokemon listesi getirir.

**ÖR:**
/api/pokemon/list?typeI=WaTer&sortBy=-baseattack

**Filtreler**
 * `typeI`
 * `typeII`

**Sıralama**
 * `name`
 * `baseAttack`

### /api/type/list
Type listesi getirir

### /api/move/list
Move listesi getirir

### /api/pokemon/:pokemonName
Ada göre bir pokemon getirir.
**ÖR:**
/api/pokemon/CHARMELEON
