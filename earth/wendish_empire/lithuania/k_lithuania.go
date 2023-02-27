package lithuania

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/courland"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/latgale"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/lithuanians"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/livonia"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/prussia"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/samogitia"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania/yatviags"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LithuaniaKingdom interface {
    feud.Kingdom
    DCourland库尔兰() 	courland.CourlandDuke
    DLatgale拉特加尔() 	latgale.LatgaleDuke
    DLithuanians立陶宛() 	lithuanians.LithuaniansDuke
    DLivonia利沃尼亚() 	livonia.LivoniaDuke
    DPrussia普鲁士() 	prussia.PrussiaDuke
    DSamogitia斯卡尔维亚() 	samogitia.SamogitiaDuke
    DYatviags约特维恩格() 	yatviags.YatviagsDuke
}

type 立陶宛LithuaniaKingdom struct {
	feud.BaseKingdom
	库尔兰Courland 	courland.CourlandDuke
	拉特加尔Latgale 	latgale.LatgaleDuke
	立陶宛Lithuanians 	lithuanians.LithuaniansDuke
	利沃尼亚Livonia 	livonia.LivoniaDuke
	普鲁士Prussia 	prussia.PrussiaDuke
	斯卡尔维亚Samogitia 	samogitia.SamogitiaDuke
	约特维恩格Yatviags 	yatviags.YatviagsDuke
}

func (k *立陶宛LithuaniaKingdom) DCourland库尔兰() courland.CourlandDuke {
	return k.库尔兰Courland
}
    
func (k *立陶宛LithuaniaKingdom) DLatgale拉特加尔() latgale.LatgaleDuke {
	return k.拉特加尔Latgale
}
    
func (k *立陶宛LithuaniaKingdom) DLithuanians立陶宛() lithuanians.LithuaniansDuke {
	return k.立陶宛Lithuanians
}
    
func (k *立陶宛LithuaniaKingdom) DLivonia利沃尼亚() livonia.LivoniaDuke {
	return k.利沃尼亚Livonia
}
    
func (k *立陶宛LithuaniaKingdom) DPrussia普鲁士() prussia.PrussiaDuke {
	return k.普鲁士Prussia
}
    
func (k *立陶宛LithuaniaKingdom) DSamogitia斯卡尔维亚() samogitia.SamogitiaDuke {
	return k.斯卡尔维亚Samogitia
}
    
func (k *立陶宛LithuaniaKingdom) DYatviags约特维恩格() yatviags.YatviagsDuke {
	return k.约特维恩格Yatviags
}
    
var KLithuania立陶宛 LithuaniaKingdom = &立陶宛LithuaniaKingdom{}

func init() {
	f := KLithuania立陶宛.(*立陶宛LithuaniaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "lithuania",
		TitleName: "立陶宛",
		TitleCode: "k_lithuania",
		Dukes:  map[string]feud.Duke{},
	}

	f.库尔兰Courland = courland.DCourland库尔兰
	f.库尔兰Courland.SetParent(f)
	
	f.拉特加尔Latgale = latgale.DLatgale拉特加尔
	f.拉特加尔Latgale.SetParent(f)
	
	f.立陶宛Lithuanians = lithuanians.DLithuanians立陶宛
	f.立陶宛Lithuanians.SetParent(f)
	
	f.利沃尼亚Livonia = livonia.DLivonia利沃尼亚
	f.利沃尼亚Livonia.SetParent(f)
	
	f.普鲁士Prussia = prussia.DPrussia普鲁士
	f.普鲁士Prussia.SetParent(f)
	
	f.斯卡尔维亚Samogitia = samogitia.DSamogitia斯卡尔维亚
	f.斯卡尔维亚Samogitia.SetParent(f)
	
	f.约特维恩格Yatviags = yatviags.DYatviags约特维恩格
	f.约特维恩格Yatviags.SetParent(f)
	
}
