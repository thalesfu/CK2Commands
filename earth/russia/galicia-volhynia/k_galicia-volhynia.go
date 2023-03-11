package galicia-volhynia

import (
	"github.com/thalesfu/CK2Commands/earth/russia/galicia-volhynia/cherven_cities"
	"github.com/thalesfu/CK2Commands/earth/russia/galicia-volhynia/galich"
	"github.com/thalesfu/CK2Commands/earth/russia/galicia-volhynia/volhynia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Galicia-volhyniaKingdom interface {
    feud.Kingdom
    DCherven_cities红鲁塞尼亚() 	cherven_cities.Cherven_citiesDuke
    DGalich加利奇() 	galich.GalichDuke
    DVolhynia沃里尼亚() 	volhynia.VolhyniaDuke
}

type 加利西亚_沃里尼亚Galicia-volhyniaKingdom struct {
	feud.BaseKingdom
	红鲁塞尼亚Cherven_cities 	cherven_cities.Cherven_citiesDuke
	加利奇Galich 	galich.GalichDuke
	沃里尼亚Volhynia 	volhynia.VolhyniaDuke
}

func (k *加利西亚_沃里尼亚Galicia-volhyniaKingdom) DCherven_cities红鲁塞尼亚() cherven_cities.Cherven_citiesDuke {
	return k.红鲁塞尼亚Cherven_cities
}
    
func (k *加利西亚_沃里尼亚Galicia-volhyniaKingdom) DGalich加利奇() galich.GalichDuke {
	return k.加利奇Galich
}
    
func (k *加利西亚_沃里尼亚Galicia-volhyniaKingdom) DVolhynia沃里尼亚() volhynia.VolhyniaDuke {
	return k.沃里尼亚Volhynia
}
    
var KGalicia-volhynia加利西亚_沃里尼亚 Galicia-volhyniaKingdom = &加利西亚_沃里尼亚Galicia-volhyniaKingdom{}

func init() {
	f := KGalicia-volhynia加利西亚_沃里尼亚.(*加利西亚_沃里尼亚Galicia-volhyniaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "galicia-volhynia",
		TitleName: "加利西亚_沃里尼亚",
		TitleCode: "k_galicia-volhynia",
		Dukes:  map[string]feud.Duke{},
	}

	f.红鲁塞尼亚Cherven_cities = cherven_cities.DCherven_cities红鲁塞尼亚
	f.红鲁塞尼亚Cherven_cities.SetParent(f)
	
	f.加利奇Galich = galich.DGalich加利奇
	f.加利奇Galich.SetParent(f)
	
	f.沃里尼亚Volhynia = volhynia.DVolhynia沃里尼亚
	f.沃里尼亚Volhynia.SetParent(f)
	
}
