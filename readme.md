### Hedefler

Groupie Trackers, bilgileri görüntüleyen bir site oluşturmak için belirli bir API'yi almayı ve içindeki verileri işlemeyi içerir.

- Dört bölümden oluşan bir [API](https://groupietrackers.herokuapp.com/api) verilecektir:

  - İlki, bazı grup ve sanatçılar hakkında isim(ler), imajı, faaliyetlerine hangi yılda başladıkları, ilk albümlerinin tarihi ve üyeleri gibi bilgileri içeren 'sanatçılar'.

  - İkincisi, `locations`, son ve/veya yaklaşan konser yerlerinden oluşur.

  - Üçüncüsü, `dates`, son ve/veya yaklaşan konser tarihlerinden oluşur.

  - Ve sonuncusu olan `relation`, diğer tüm parçalar, yani "sanatçılar", "tarihler" ve "yerler" arasındaki bağlantıyı kurar.

- Tüm bunları göz önünde bulundurarak, çeşitli veri görselleştirmeleri (örnekler: bloklar, kartlar, tablolar, liste, sayfalar, grafikler vb.) aracılığıyla bant bilgilerini görüntüleyebileceğiniz kullanıcı dostu bir web sitesi oluşturmalısınız. Nasıl sergileyeceğinize karar vermek size kalmıştır.

- Bu proje aynı zamanda olayların/eylemlerin yaratılmasına ve görselleştirilmesine de odaklanmaktadır.

  - Yapmanızı istediğimiz olay/eylem, sunucuya (istemci-sunucu) istemci çağrısı olarak bilinir. Bir eylemi tetiklemesi gereken, sizin seçeceğiniz bir özellik diyebiliriz. Bu eylemin bilgi alabilmesi için sunucuyla iletişim kurması gerekir, ([request-response])(https://en.wikipedia.org/wiki/Request%E2%80%93response)
  - Bir olay, müşteri, zaman veya başka bir faktör tarafından tetiklenen bir tür eyleme yanıt veren bir sistemden oluşur.

### Talimatlar

- Arka uç **Go** ile yazılmalıdır.
- Site ve sunucu herhangi bir zamanda çökemez.
- Tüm sayfalar düzgün çalışmalı ve olası hatalara dikkat etmelisiniz. 
- Kod, [**iyi uygulamalar**](../good-practices/README.md) kurallarına uygun olmalıdır.
- [unit testing](https://go.dev/doc/tutorial/add-a-test) için **test dosyalarının** olması önerilir.

### Kullanım

- Bir RESTful API örneğini [here](https://rickandmortyapi.com/) görebilirsiniz.

Bu proje aşağıdakileri öğrenmenize yardımcı olacaktır:

- Verilerin işlenmesi ve saklanması.
- [JSON](https://www.json.org/json-en.html) dosyaları ve biçimi.
-HTML.
- Etkinlik oluşturma ve görüntüleme.
- [Client-server](https://developer.mozilla.org/en-US/docs/Learn/Server-side/First_steps/Client-Server_overview).