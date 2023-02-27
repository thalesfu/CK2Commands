package aquitaine

import (
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/aquitaine"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/auvergne"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/bourbon"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/gascogne"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/poitou"
	"github.com/thalesfu/CK2Commands/earth/france/aquitaine/toulouse"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AquitaineKingdom interface {
    feud.Kingdom
    DAquitaine阿基坦() 	aquitaine.AquitaineDuke
    DAuvergne奥弗涅() 	auvergne.AuvergneDuke
    DBourbon波旁() 	bourbon.BourbonDuke
    DGascogne加斯科涅() 	gascogne.GascogneDuke
    DPoitou普瓦图() 	poitou.PoitouDuke
    DToulouse图卢兹() 	toulouse.ToulouseDuke
}

type 阿基坦AquitaineKingdom struct {
	feud.BaseKingdom
	阿基坦Aquitaine 	aquitaine.AquitaineDuke
	奥弗涅Auvergne 	auvergne.AuvergneDuke
	波旁Bourbon 	bourbon.BourbonDuke
	加斯科涅Gascogne 	gascogne.GascogneDuke
	普瓦图Poitou 	poitou.PoitouDuke
	图卢兹Toulouse 	toulouse.ToulouseDuke
}

func (k *阿基坦AquitaineKingdom) DAquitaine阿基坦() aquitaine.AquitaineDuke {
	return k.阿基坦Aquitaine
}
    
func (k *阿基坦AquitaineKingdom) DAuvergne奥弗涅() auvergne.AuvergneDuke {
	return k.奥弗涅Auvergne
}
    
func (k *阿基坦AquitaineKingdom) DBourbon波旁() bourbon.BourbonDuke {
	return k.波旁Bourbon
}
    
func (k *阿基坦AquitaineKingdom) DGascogne加斯科涅() gascogne.GascogneDuke {
	return k.加斯科涅Gascogne
}
    
func (k *阿基坦AquitaineKingdom) DPoitou普瓦图() poitou.PoitouDuke {
	return k.普瓦图Poitou
}
    
func (k *阿基坦AquitaineKingdom) DToulouse图卢兹() toulouse.ToulouseDuke {
	return k.图卢兹Toulouse
}
    
var KAquitaine阿基坦 AquitaineKingdom = &阿基坦AquitaineKingdom{}

func init() {
	f := KAquitaine阿基坦.(*阿基坦AquitaineKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "aquitaine",
		TitleName: "阿基坦",
		TitleCode: "k_aquitaine",
		Dukes:  map[string]feud.Duke{},
	}

	f.阿基坦Aquitaine = aquitaine.DAquitaine阿基坦
	f.阿基坦Aquitaine.SetParent(f)
	
	f.奥弗涅Auvergne = auvergne.DAuvergne奥弗涅
	f.奥弗涅Auvergne.SetParent(f)
	
	f.波旁Bourbon = bourbon.DBourbon波旁
	f.波旁Bourbon.SetParent(f)
	
	f.加斯科涅Gascogne = gascogne.DGascogne加斯科涅
	f.加斯科涅Gascogne.SetParent(f)
	
	f.普瓦图Poitou = poitou.DPoitou普瓦图
	f.普瓦图Poitou.SetParent(f)
	
	f.图卢兹Toulouse = toulouse.DToulouse图卢兹
	f.图卢兹Toulouse.SetParent(f)
	
}
