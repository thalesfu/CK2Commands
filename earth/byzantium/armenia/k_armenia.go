package armenia

import (
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/armenia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/coloneia"
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/edessa"
	"github.com/thalesfu/CK2Commands/earth/byzantium/armenia/mesopotamia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ArmeniaKingdom interface {
    feud.Kingdom
    DArmenia亚美尼亚() 	armenia.ArmeniaDuke
    DColoneia科洛尼亚() 	coloneia.ColoneiaDuke
    DEdessa埃德萨() 	edessa.EdessaDuke
    DMesopotamia美索不达米亚() 	mesopotamia.MesopotamiaDuke
}

type 亚美尼亚ArmeniaKingdom struct {
	feud.BaseKingdom
	亚美尼亚Armenia 	armenia.ArmeniaDuke
	科洛尼亚Coloneia 	coloneia.ColoneiaDuke
	埃德萨Edessa 	edessa.EdessaDuke
	美索不达米亚Mesopotamia 	mesopotamia.MesopotamiaDuke
}

func (k *亚美尼亚ArmeniaKingdom) DArmenia亚美尼亚() armenia.ArmeniaDuke {
	return k.亚美尼亚Armenia
}
    
func (k *亚美尼亚ArmeniaKingdom) DColoneia科洛尼亚() coloneia.ColoneiaDuke {
	return k.科洛尼亚Coloneia
}
    
func (k *亚美尼亚ArmeniaKingdom) DEdessa埃德萨() edessa.EdessaDuke {
	return k.埃德萨Edessa
}
    
func (k *亚美尼亚ArmeniaKingdom) DMesopotamia美索不达米亚() mesopotamia.MesopotamiaDuke {
	return k.美索不达米亚Mesopotamia
}
    
var KArmenia亚美尼亚 ArmeniaKingdom = &亚美尼亚ArmeniaKingdom{}

func init() {
	f := KArmenia亚美尼亚.(*亚美尼亚ArmeniaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "armenia",
		TitleName: "亚美尼亚",
		TitleCode: "k_armenia",
		Dukes:  map[string]feud.Duke{},
	}

	f.亚美尼亚Armenia = armenia.DArmenia亚美尼亚
	f.亚美尼亚Armenia.SetParent(f)
	
	f.科洛尼亚Coloneia = coloneia.DColoneia科洛尼亚
	f.科洛尼亚Coloneia.SetParent(f)
	
	f.埃德萨Edessa = edessa.DEdessa埃德萨
	f.埃德萨Edessa.SetParent(f)
	
	f.美索不达米亚Mesopotamia = mesopotamia.DMesopotamia美索不达米亚
	f.美索不达米亚Mesopotamia.SetParent(f)
	
}
