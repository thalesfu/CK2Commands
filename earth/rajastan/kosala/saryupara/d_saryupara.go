package saryupara

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/saryupara/ayodhya"
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/saryupara/naimisa"
	"github.com/thalesfu/CK2Commands/earth/rajastan/kosala/saryupara/sravasti"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaryuparaDuke interface {
    feud.Duke
    CAyodhya阿踰陀() 	ayodhya.AyodhyaCounty
    CNaimisa泥弥沙() 	naimisa.NaimisaCounty
    CSravasti舍卫城() 	sravasti.SravastiCounty
}

type 萨罗瑜波罗SaryuparaDuke struct {
	feud.BaseDuke
	阿踰陀Ayodhya 	ayodhya.AyodhyaCounty
	泥弥沙Naimisa 	naimisa.NaimisaCounty
	舍卫城Sravasti 	sravasti.SravastiCounty
}

func (d *萨罗瑜波罗SaryuparaDuke) CAyodhya阿踰陀() ayodhya.AyodhyaCounty {
	return d.阿踰陀Ayodhya
}
    
func (d *萨罗瑜波罗SaryuparaDuke) CNaimisa泥弥沙() naimisa.NaimisaCounty {
	return d.泥弥沙Naimisa
}
    
func (d *萨罗瑜波罗SaryuparaDuke) CSravasti舍卫城() sravasti.SravastiCounty {
	return d.舍卫城Sravasti
}
    
var DSaryupara萨罗瑜波罗 SaryuparaDuke = &萨罗瑜波罗SaryuparaDuke{}

func init() {
	f := DSaryupara萨罗瑜波罗.(*萨罗瑜波罗SaryuparaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "saryupara",
		TitleName: "萨罗瑜波罗",
		TitleCode: "d_saryupara",
		Counties:  map[string]feud.County{},
	}

	f.阿踰陀Ayodhya = ayodhya.CAyodhya阿踰陀
	f.阿踰陀Ayodhya.SetParent(f)
	
	f.泥弥沙Naimisa = naimisa.CNaimisa泥弥沙
	f.泥弥沙Naimisa.SetParent(f)
	
	f.舍卫城Sravasti = sravasti.CSravasti舍卫城
	f.舍卫城Sravasti.SetParent(f)
	
}
