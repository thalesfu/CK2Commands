package sistan

import (
	"github.com/thalesfu/CK2Commands/earth/persia/baluchistan/sistan/bost"
	"github.com/thalesfu/CK2Commands/earth/persia/baluchistan/sistan/farrah"
	"github.com/thalesfu/CK2Commands/earth/persia/baluchistan/sistan/zahedan"
	"github.com/thalesfu/CK2Commands/earth/persia/baluchistan/sistan/zaranj"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SistanDuke interface {
    feud.Duke
    CBost博斯特() 	bost.BostCounty
    CFarrah法拉() 	farrah.FarrahCounty
    CZahedan扎黑丹() 	zahedan.ZahedanCounty
    CZaranj疾陵() 	zaranj.ZaranjCounty
}

type 昔思丹SistanDuke struct {
	feud.BaseDuke
	博斯特Bost 	bost.BostCounty
	法拉Farrah 	farrah.FarrahCounty
	扎黑丹Zahedan 	zahedan.ZahedanCounty
	疾陵Zaranj 	zaranj.ZaranjCounty
}

func (d *昔思丹SistanDuke) CBost博斯特() bost.BostCounty {
	return d.博斯特Bost
}
    
func (d *昔思丹SistanDuke) CFarrah法拉() farrah.FarrahCounty {
	return d.法拉Farrah
}
    
func (d *昔思丹SistanDuke) CZahedan扎黑丹() zahedan.ZahedanCounty {
	return d.扎黑丹Zahedan
}
    
func (d *昔思丹SistanDuke) CZaranj疾陵() zaranj.ZaranjCounty {
	return d.疾陵Zaranj
}
    
var DSistan昔思丹 SistanDuke = &昔思丹SistanDuke{}

func init() {
	f := DSistan昔思丹.(*昔思丹SistanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "sistan",
		TitleName: "昔思丹",
		TitleCode: "d_sistan",
		Counties:  map[string]feud.County{},
	}

	f.博斯特Bost = bost.CBost博斯特
	f.博斯特Bost.SetParent(f)
	
	f.法拉Farrah = farrah.CFarrah法拉
	f.法拉Farrah.SetParent(f)
	
	f.扎黑丹Zahedan = zahedan.CZahedan扎黑丹
	f.扎黑丹Zahedan.SetParent(f)
	
	f.疾陵Zaranj = zaranj.CZaranj疾陵
	f.疾陵Zaranj.SetParent(f)
	
}
