package utils

import (
	"github.com/erenerdogmus/dataloaders"
	"github.com/erenerdogmus/models"
	"strconv"
	"strings"
	"sync"
)

var (
	allData     []models.Data // Tüm sanatçı verilerinin önbelleğe alındığı global değişken
	allDataOnce sync.Once     // Tüm verilerin bir kez toplanıp önbelleğe alınmasını sağlar
)

// GetAllData, tüm sanatçı verilerini toplar ve bunları önbellekte saklar.
// Bu işlev, verileri yalnızca ilk çağrıldığında toplar ve sonraki çağrılarda önbellekten döndürür.
func GetAllData() []models.Data {
	// Tüm veriler bir kez alındıysa tekrar alınmaması için kontrol edilir.
	allDataOnce.Do(func() {
		artistData := dataloaders.ArtistData() // Sanatçı verilerini yükler
		allData = make([]models.Data, len(artistData))
		var wg sync.WaitGroup
		wg.Add(len(artistData))
		for i, artist := range artistData {
			go func(index int, id int) {
				defer wg.Done()
				allData[index] = CollectSingleData(strconv.Itoa(id)) // Her sanatçı için veri toplama işlemi
			}(i, artist.Id)
		}
		wg.Wait() // Tüm goroutine'lerin bitmesini bekler
	})
	return allData
}

// CollectData, verilen arama terimine göre sanatçı verilerini filtreler.
// Arama terimi boşsa tüm verileri döndürür, değilse terimi içeren verileri filtreleyip döndürür.
func CollectData(searchTerm string) []models.Data {
	if searchTerm == "" {
		return GetAllData()
	}
	var matchingDatas []models.Data
	allData := GetAllData()

	var wg sync.WaitGroup
	wg.Add(len(allData))
	for _, data := range allData {
		go func(d models.Data) {
			defer wg.Done()
			if ContainsMatch(d, searchTerm) {
				matchingDatas = append(matchingDatas, d) // Eşleşme bulunan verileri listeye ekler
			}
		}(data)
	}
	wg.Wait() // Tüm goroutine'lerin bitmesini bekler
	return matchingDatas
}

// ContainsMatch, verilen veri ve arama terimi ile eşleşme olup olmadığını kontrol eder.
// Eşleşme; sanatçı adı, albüm ismi, oluşturulma tarihi, üyeler ve konumlar üzerinden yapılır.
func ContainsMatch(data models.Data, term string) bool {
	if strings.Contains(strings.ToUpper(data.A.Name), term) ||
		strings.Contains(data.A.FirstAlbum, term) ||
		strings.Contains(strconv.Itoa(data.A.CreationDate), term) {
		return true
	}
	for _, members := range data.A.Members {
		if strings.Contains(strings.ToUpper(members), term) {
			return true
		}
	}
	for _, location := range data.L.Locations {
		if strings.Contains(strings.ToUpper(location), term) {
			return true
		}
	}
	return false
}

// CollectSingleData, belirli bir ID'ye sahip sanatçı için tüm ilgili verileri toplar.
// Toplanan veriler sanatçı bilgileri, ilişkiler, tarihler ve konumları içerir.
func CollectSingleData(id string) models.Data {
	result := models.Data{
		A: GetArtistID(id),          // Sanatçı bilgilerini yükler
		R: dataloaders.Relation(id), // İlişkileri yükler
		D: dataloaders.Dates(id),    // Tarih bilgilerini yükler
		L: dataloaders.Location(id), // Konum bilgilerini yükler
	}
	return result
}

// GetArtistID, belirli bir ID'ye sahip sanatçının bilgilerini döndürür.
// Eğer sanatçı bulunursa, sanatçının bilgileri döndürülür.
func GetArtistID(id string) models.ArtistStruct {
	var result models.ArtistStruct
	for _, artist := range dataloaders.ArtistData() {
		if strconv.Itoa(artist.Id) == id {
			result = artist
			break // Eşleşen ID bulunduğunda döngüden çık
		}
	}
	return result
}
