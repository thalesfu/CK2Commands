package estonia

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/estonia/esthonia"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/estonia/sakala"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EstoniaKingdom interface {
    feud.Kingdom
    DEsthonia爱沙尼亚() 	esthonia.EsthoniaDuke
    DSakala萨卡拉() 	sakala.SakalaDuke
}

type 爱沙尼亚EstoniaKingdom struct {
	feud.BaseKingdom
	爱沙尼亚Esthonia 	esthonia.EsthoniaDuke
	萨卡拉Sakala 	sakala.SakalaDuke
}

func (k *爱沙尼亚EstoniaKingdom) DEsthonia爱沙尼亚() esthonia.EsthoniaDuke {
	return k.爱沙尼亚Esthonia
}
    
func (k *爱沙尼亚EstoniaKingdom) DSakala萨卡拉() sakala.SakalaDuke {
	return k.萨卡拉Sakala
}
    
var KEstonia爱沙尼亚 EstoniaKingdom = &爱沙尼亚EstoniaKingdom{}

func init() {
	f := KEstonia爱沙尼亚.(*爱沙尼亚EstoniaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "estonia",
		TitleName: "爱沙尼亚",
		TitleCode: "k_estonia",
		Dukes:  map[string]feud.Duke{},
	}

	f.爱沙尼亚Esthonia = esthonia.DEsthonia爱沙尼亚
	f.爱沙尼亚Esthonia.SetParent(f)
	
	f.萨卡拉Sakala = sakala.DSakala萨卡拉
	f.萨卡拉Sakala.SetParent(f)
	
}
