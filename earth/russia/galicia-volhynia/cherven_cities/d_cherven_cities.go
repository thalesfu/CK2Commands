package cherven_cities

import (
	"github.com/thalesfu/CK2Commands/earth/russia/galicia-volhynia/cherven_cities/belz"
	"github.com/thalesfu/CK2Commands/earth/russia/galicia-volhynia/cherven_cities/cherven"
	"github.com/thalesfu/CK2Commands/earth/russia/galicia-volhynia/cherven_cities/peremyshl"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Cherven_citiesDuke interface {
    feud.Duke
    CBelz贝尔兹() 	belz.BelzCounty
    CCherven切尔文() 	cherven.ChervenCounty
    CPeremyshl佩列梅什利() 	peremyshl.PeremyshlCounty
}

type 红鲁塞尼亚Cherven_citiesDuke struct {
	feud.BaseDuke
	贝尔兹Belz 	belz.BelzCounty
	切尔文Cherven 	cherven.ChervenCounty
	佩列梅什利Peremyshl 	peremyshl.PeremyshlCounty
}

func (d *红鲁塞尼亚Cherven_citiesDuke) CBelz贝尔兹() belz.BelzCounty {
	return d.贝尔兹Belz
}
    
func (d *红鲁塞尼亚Cherven_citiesDuke) CCherven切尔文() cherven.ChervenCounty {
	return d.切尔文Cherven
}
    
func (d *红鲁塞尼亚Cherven_citiesDuke) CPeremyshl佩列梅什利() peremyshl.PeremyshlCounty {
	return d.佩列梅什利Peremyshl
}
    
var DCherven_cities红鲁塞尼亚 Cherven_citiesDuke = &红鲁塞尼亚Cherven_citiesDuke{}

func init() {
	f := DCherven_cities红鲁塞尼亚.(*红鲁塞尼亚Cherven_citiesDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "cherven_cities",
		TitleName: "红鲁塞尼亚",
		TitleCode: "d_cherven_cities",
		Counties:  map[string]feud.County{},
	}

	f.贝尔兹Belz = belz.CBelz贝尔兹
	f.贝尔兹Belz.SetParent(f)
	
	f.切尔文Cherven = cherven.CCherven切尔文
	f.切尔文Cherven.SetParent(f)
	
	f.佩列梅什利Peremyshl = peremyshl.CPeremyshl佩列梅什利
	f.佩列梅什利Peremyshl.SetParent(f)
	
}
