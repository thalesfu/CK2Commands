package aragon

import (
	"github.com/thalesfu/CK2Commands/earth/spain/aragon/aragon"
	"github.com/thalesfu/CK2Commands/earth/spain/aragon/barcelona"
	"github.com/thalesfu/CK2Commands/earth/spain/aragon/mallorca"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AragonKingdom interface {
    feud.Kingdom
    DAragon阿拉贡() 	aragon.AragonDuke
    DBarcelona巴塞罗那() 	barcelona.BarcelonaDuke
    DMallorca马略卡() 	mallorca.MallorcaDuke
}

type 阿拉贡AragonKingdom struct {
	feud.BaseKingdom
	阿拉贡Aragon 	aragon.AragonDuke
	巴塞罗那Barcelona 	barcelona.BarcelonaDuke
	马略卡Mallorca 	mallorca.MallorcaDuke
}

func (k *阿拉贡AragonKingdom) DAragon阿拉贡() aragon.AragonDuke {
	return k.阿拉贡Aragon
}
    
func (k *阿拉贡AragonKingdom) DBarcelona巴塞罗那() barcelona.BarcelonaDuke {
	return k.巴塞罗那Barcelona
}
    
func (k *阿拉贡AragonKingdom) DMallorca马略卡() mallorca.MallorcaDuke {
	return k.马略卡Mallorca
}
    
var KAragon阿拉贡 AragonKingdom = &阿拉贡AragonKingdom{}

func init() {
	f := KAragon阿拉贡.(*阿拉贡AragonKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "aragon",
		TitleName: "阿拉贡",
		TitleCode: "k_aragon",
		Dukes:  map[string]feud.Duke{},
	}

	f.阿拉贡Aragon = aragon.DAragon阿拉贡
	f.阿拉贡Aragon.SetParent(f)
	
	f.巴塞罗那Barcelona = barcelona.DBarcelona巴塞罗那
	f.巴塞罗那Barcelona.SetParent(f)
	
	f.马略卡Mallorca = mallorca.DMallorca马略卡
	f.马略卡Mallorca.SetParent(f)
	
}
