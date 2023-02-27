package nenets

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/bjarmia"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/hlynov"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/kargopol"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/pechora"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/samoyed"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/zavarot"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NenetsKingdom interface {
    feud.Kingdom
    DBjarmia比亚尔米亚() 	bjarmia.BjarmiaDuke
    DHlynov韦利斯克() 	hlynov.HlynovDuke
    DKargopol卡尔戈波尔() 	kargopol.KargopolDuke
    DPechora伯朝拉() 	pechora.PechoraDuke
    DSamoyed梅津() 	samoyed.SamoyedDuke
    DZavarot扎瓦罗特() 	zavarot.ZavarotDuke
}

type 涅涅茨NenetsKingdom struct {
	feud.BaseKingdom
	比亚尔米亚Bjarmia 	bjarmia.BjarmiaDuke
	韦利斯克Hlynov 	hlynov.HlynovDuke
	卡尔戈波尔Kargopol 	kargopol.KargopolDuke
	伯朝拉Pechora 	pechora.PechoraDuke
	梅津Samoyed 	samoyed.SamoyedDuke
	扎瓦罗特Zavarot 	zavarot.ZavarotDuke
}

func (k *涅涅茨NenetsKingdom) DBjarmia比亚尔米亚() bjarmia.BjarmiaDuke {
	return k.比亚尔米亚Bjarmia
}
    
func (k *涅涅茨NenetsKingdom) DHlynov韦利斯克() hlynov.HlynovDuke {
	return k.韦利斯克Hlynov
}
    
func (k *涅涅茨NenetsKingdom) DKargopol卡尔戈波尔() kargopol.KargopolDuke {
	return k.卡尔戈波尔Kargopol
}
    
func (k *涅涅茨NenetsKingdom) DPechora伯朝拉() pechora.PechoraDuke {
	return k.伯朝拉Pechora
}
    
func (k *涅涅茨NenetsKingdom) DSamoyed梅津() samoyed.SamoyedDuke {
	return k.梅津Samoyed
}
    
func (k *涅涅茨NenetsKingdom) DZavarot扎瓦罗特() zavarot.ZavarotDuke {
	return k.扎瓦罗特Zavarot
}
    
var KNenets涅涅茨 NenetsKingdom = &涅涅茨NenetsKingdom{}

func init() {
	f := KNenets涅涅茨.(*涅涅茨NenetsKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "nenets",
		TitleName: "涅涅茨",
		TitleCode: "k_nenets",
		Dukes:  map[string]feud.Duke{},
	}

	f.比亚尔米亚Bjarmia = bjarmia.DBjarmia比亚尔米亚
	f.比亚尔米亚Bjarmia.SetParent(f)
	
	f.韦利斯克Hlynov = hlynov.DHlynov韦利斯克
	f.韦利斯克Hlynov.SetParent(f)
	
	f.卡尔戈波尔Kargopol = kargopol.DKargopol卡尔戈波尔
	f.卡尔戈波尔Kargopol.SetParent(f)
	
	f.伯朝拉Pechora = pechora.DPechora伯朝拉
	f.伯朝拉Pechora.SetParent(f)
	
	f.梅津Samoyed = samoyed.DSamoyed梅津
	f.梅津Samoyed.SetParent(f)
	
	f.扎瓦罗特Zavarot = zavarot.DZavarot扎瓦罗特
	f.扎瓦罗特Zavarot.SetParent(f)
	
}
