package volhynia

import (
	"github.com/thalesfu/CK2Commands/earth/russia/galicia-volhynia/volhynia/beresty"
	"github.com/thalesfu/CK2Commands/earth/russia/galicia-volhynia/volhynia/lutsk"
	"github.com/thalesfu/CK2Commands/earth/russia/galicia-volhynia/volhynia/podlasie"
	"github.com/thalesfu/CK2Commands/earth/russia/galicia-volhynia/volhynia/vladimir_volynsky"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VolhyniaDuke interface {
    feud.Duke
    CBeresty别列斯捷() 	beresty.BerestyCounty
    CLutsk卢茨克() 	lutsk.LutskCounty
    CPodlasie波德拉谢() 	podlasie.PodlasieCounty
    CVladimir_volynsky弗拉基米尔沃伦斯基() 	vladimir_volynsky.Vladimir_volynskyCounty
}

type 沃里尼亚VolhyniaDuke struct {
	feud.BaseDuke
	别列斯捷Beresty 	beresty.BerestyCounty
	卢茨克Lutsk 	lutsk.LutskCounty
	波德拉谢Podlasie 	podlasie.PodlasieCounty
	弗拉基米尔沃伦斯基Vladimir_volynsky 	vladimir_volynsky.Vladimir_volynskyCounty
}

func (d *沃里尼亚VolhyniaDuke) CBeresty别列斯捷() beresty.BerestyCounty {
	return d.别列斯捷Beresty
}
    
func (d *沃里尼亚VolhyniaDuke) CLutsk卢茨克() lutsk.LutskCounty {
	return d.卢茨克Lutsk
}
    
func (d *沃里尼亚VolhyniaDuke) CPodlasie波德拉谢() podlasie.PodlasieCounty {
	return d.波德拉谢Podlasie
}
    
func (d *沃里尼亚VolhyniaDuke) CVladimir_volynsky弗拉基米尔沃伦斯基() vladimir_volynsky.Vladimir_volynskyCounty {
	return d.弗拉基米尔沃伦斯基Vladimir_volynsky
}
    
var DVolhynia沃里尼亚 VolhyniaDuke = &沃里尼亚VolhyniaDuke{}

func init() {
	f := DVolhynia沃里尼亚.(*沃里尼亚VolhyniaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "volhynia",
		TitleName: "沃里尼亚",
		TitleCode: "d_volhynia",
		Counties:  map[string]feud.County{},
	}

	f.别列斯捷Beresty = beresty.CBeresty别列斯捷
	f.别列斯捷Beresty.SetParent(f)
	
	f.卢茨克Lutsk = lutsk.CLutsk卢茨克
	f.卢茨克Lutsk.SetParent(f)
	
	f.波德拉谢Podlasie = podlasie.CPodlasie波德拉谢
	f.波德拉谢Podlasie.SetParent(f)
	
	f.弗拉基米尔沃伦斯基Vladimir_volynsky = vladimir_volynsky.CVladimir_volynsky弗拉基米尔沃伦斯基
	f.弗拉基米尔沃伦斯基Vladimir_volynsky.SetParent(f)
	
}
