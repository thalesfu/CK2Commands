package thrace

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/abydos"
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/adrianopolis"
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/optimatoi"
	"github.com/thalesfu/CK2Commands/earth/byzantium/thrace/thrace"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ThraceKingdom interface {
    feud.Kingdom
    DAbydos阿卑多斯() 	abydos.AbydosDuke
    DAdrianopolis哈德良波利斯() 	adrianopolis.AdrianopolisDuke
    DOptimatoi奥普提马通() 	optimatoi.OptimatoiDuke
    DThrace色雷斯() 	thrace.ThraceDuke
}

type 色雷斯ThraceKingdom struct {
	feud.BaseKingdom
	阿卑多斯Abydos 	abydos.AbydosDuke
	哈德良波利斯Adrianopolis 	adrianopolis.AdrianopolisDuke
	奥普提马通Optimatoi 	optimatoi.OptimatoiDuke
	色雷斯Thrace 	thrace.ThraceDuke
}

func (k *色雷斯ThraceKingdom) DAbydos阿卑多斯() abydos.AbydosDuke {
	return k.阿卑多斯Abydos
}
    
func (k *色雷斯ThraceKingdom) DAdrianopolis哈德良波利斯() adrianopolis.AdrianopolisDuke {
	return k.哈德良波利斯Adrianopolis
}
    
func (k *色雷斯ThraceKingdom) DOptimatoi奥普提马通() optimatoi.OptimatoiDuke {
	return k.奥普提马通Optimatoi
}
    
func (k *色雷斯ThraceKingdom) DThrace色雷斯() thrace.ThraceDuke {
	return k.色雷斯Thrace
}
    
var KThrace色雷斯 ThraceKingdom = &色雷斯ThraceKingdom{}

func init() {
	f := KThrace色雷斯.(*色雷斯ThraceKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "thrace",
		TitleName: "色雷斯",
		TitleCode: "k_thrace",
		Dukes:  map[string]feud.Duke{},
	}

	f.阿卑多斯Abydos = abydos.DAbydos阿卑多斯
	f.阿卑多斯Abydos.SetParent(f)
	
	f.哈德良波利斯Adrianopolis = adrianopolis.DAdrianopolis哈德良波利斯
	f.哈德良波利斯Adrianopolis.SetParent(f)
	
	f.奥普提马通Optimatoi = optimatoi.DOptimatoi奥普提马通
	f.奥普提马通Optimatoi.SetParent(f)
	
	f.色雷斯Thrace = thrace.DThrace色雷斯
	f.色雷斯Thrace.SetParent(f)
	
}
