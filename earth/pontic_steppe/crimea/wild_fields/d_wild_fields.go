package wild_fields

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/wild_fields/bratslav"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/wild_fields/korsun"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/wild_fields/odessa"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/wild_fields/oleshye"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/wild_fields/olvia"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/wild_fields/vinnytsia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Wild_fieldsDuke interface {
    feud.Duke
    CBratslav布拉茨拉夫() 	bratslav.BratslavCounty
    CKorsun科尔孙() 	korsun.KorsunCounty
    COdessa敖德萨() 	odessa.OdessaCounty
    COleshye下第聂伯() 	oleshye.OleshyeCounty
    COlvia奥利维亚() 	olvia.OlviaCounty
    CVinnytsia文尼察() 	vinnytsia.VinnytsiaCounty
}

type 艾泰尔克兹Wild_fieldsDuke struct {
	feud.BaseDuke
	布拉茨拉夫Bratslav 	bratslav.BratslavCounty
	科尔孙Korsun 	korsun.KorsunCounty
	敖德萨Odessa 	odessa.OdessaCounty
	下第聂伯Oleshye 	oleshye.OleshyeCounty
	奥利维亚Olvia 	olvia.OlviaCounty
	文尼察Vinnytsia 	vinnytsia.VinnytsiaCounty
}

func (d *艾泰尔克兹Wild_fieldsDuke) CBratslav布拉茨拉夫() bratslav.BratslavCounty {
	return d.布拉茨拉夫Bratslav
}
    
func (d *艾泰尔克兹Wild_fieldsDuke) CKorsun科尔孙() korsun.KorsunCounty {
	return d.科尔孙Korsun
}
    
func (d *艾泰尔克兹Wild_fieldsDuke) COdessa敖德萨() odessa.OdessaCounty {
	return d.敖德萨Odessa
}
    
func (d *艾泰尔克兹Wild_fieldsDuke) COleshye下第聂伯() oleshye.OleshyeCounty {
	return d.下第聂伯Oleshye
}
    
func (d *艾泰尔克兹Wild_fieldsDuke) COlvia奥利维亚() olvia.OlviaCounty {
	return d.奥利维亚Olvia
}
    
func (d *艾泰尔克兹Wild_fieldsDuke) CVinnytsia文尼察() vinnytsia.VinnytsiaCounty {
	return d.文尼察Vinnytsia
}
    
var DWild_fields艾泰尔克兹 Wild_fieldsDuke = &艾泰尔克兹Wild_fieldsDuke{}

func init() {
	f := DWild_fields艾泰尔克兹.(*艾泰尔克兹Wild_fieldsDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "wild_fields",
		TitleName: "艾泰尔克兹",
		TitleCode: "d_wild_fields",
		Counties:  map[string]feud.County{},
	}

	f.布拉茨拉夫Bratslav = bratslav.CBratslav布拉茨拉夫
	f.布拉茨拉夫Bratslav.SetParent(f)
	
	f.科尔孙Korsun = korsun.CKorsun科尔孙
	f.科尔孙Korsun.SetParent(f)
	
	f.敖德萨Odessa = odessa.COdessa敖德萨
	f.敖德萨Odessa.SetParent(f)
	
	f.下第聂伯Oleshye = oleshye.COleshye下第聂伯
	f.下第聂伯Oleshye.SetParent(f)
	
	f.奥利维亚Olvia = olvia.COlvia奥利维亚
	f.奥利维亚Olvia.SetParent(f)
	
	f.文尼察Vinnytsia = vinnytsia.CVinnytsia文尼察
	f.文尼察Vinnytsia.SetParent(f)
	
}
