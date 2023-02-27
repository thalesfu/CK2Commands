package usturt

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/usturt/buzachi"
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/usturt/kara_bogaz"
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/usturt/kusbulak"
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/usturt/mangyshlak"
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/usturt/usturt"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UsturtDuke interface {
    feud.Duke
    CBuzachi布扎奇() 	buzachi.BuzachiCounty
    CKara_bogaz卡拉博加兹() 	kara_bogaz.Kara_bogazCounty
    CKusbulak库斯布拉克() 	kusbulak.KusbulakCounty
    CMangyshlak曼格什拉克() 	mangyshlak.MangyshlakCounty
    CUsturt乌斯秋尔特() 	usturt.UsturtCounty
}

type 乌斯秋尔特UsturtDuke struct {
	feud.BaseDuke
	布扎奇Buzachi 	buzachi.BuzachiCounty
	卡拉博加兹Kara_bogaz 	kara_bogaz.Kara_bogazCounty
	库斯布拉克Kusbulak 	kusbulak.KusbulakCounty
	曼格什拉克Mangyshlak 	mangyshlak.MangyshlakCounty
	乌斯秋尔特Usturt 	usturt.UsturtCounty
}

func (d *乌斯秋尔特UsturtDuke) CBuzachi布扎奇() buzachi.BuzachiCounty {
	return d.布扎奇Buzachi
}
    
func (d *乌斯秋尔特UsturtDuke) CKara_bogaz卡拉博加兹() kara_bogaz.Kara_bogazCounty {
	return d.卡拉博加兹Kara_bogaz
}
    
func (d *乌斯秋尔特UsturtDuke) CKusbulak库斯布拉克() kusbulak.KusbulakCounty {
	return d.库斯布拉克Kusbulak
}
    
func (d *乌斯秋尔特UsturtDuke) CMangyshlak曼格什拉克() mangyshlak.MangyshlakCounty {
	return d.曼格什拉克Mangyshlak
}
    
func (d *乌斯秋尔特UsturtDuke) CUsturt乌斯秋尔特() usturt.UsturtCounty {
	return d.乌斯秋尔特Usturt
}
    
var DUsturt乌斯秋尔特 UsturtDuke = &乌斯秋尔特UsturtDuke{}

func init() {
	f := DUsturt乌斯秋尔特.(*乌斯秋尔特UsturtDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "usturt",
		TitleName: "乌斯秋尔特",
		TitleCode: "d_usturt",
		Counties:  map[string]feud.County{},
	}

	f.布扎奇Buzachi = buzachi.CBuzachi布扎奇
	f.布扎奇Buzachi.SetParent(f)
	
	f.卡拉博加兹Kara_bogaz = kara_bogaz.CKara_bogaz卡拉博加兹
	f.卡拉博加兹Kara_bogaz.SetParent(f)
	
	f.库斯布拉克Kusbulak = kusbulak.CKusbulak库斯布拉克
	f.库斯布拉克Kusbulak.SetParent(f)
	
	f.曼格什拉克Mangyshlak = mangyshlak.CMangyshlak曼格什拉克
	f.曼格什拉克Mangyshlak.SetParent(f)
	
	f.乌斯秋尔特Usturt = usturt.CUsturt乌斯秋尔特
	f.乌斯秋尔特Usturt.SetParent(f)
	
}
