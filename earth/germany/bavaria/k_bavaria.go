package bavaria

import (
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/bavaria"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/nordgau"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/osterreich"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/salzburg"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/styria"
	"github.com/thalesfu/CK2Commands/earth/germany/bavaria/tyrol"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BavariaKingdom interface {
    feud.Kingdom
    DBavaria巴伐利亚() 	bavaria.BavariaDuke
    DNordgau诺德高() 	nordgau.NordgauDuke
    DOsterreich奥地利() 	osterreich.OsterreichDuke
    DSalzburg萨尔茨堡() 	salzburg.SalzburgDuke
    DStyria施蒂里亚() 	styria.StyriaDuke
    DTyrol蒂罗尔() 	tyrol.TyrolDuke
}

type 巴伐利亚BavariaKingdom struct {
	feud.BaseKingdom
	巴伐利亚Bavaria 	bavaria.BavariaDuke
	诺德高Nordgau 	nordgau.NordgauDuke
	奥地利Osterreich 	osterreich.OsterreichDuke
	萨尔茨堡Salzburg 	salzburg.SalzburgDuke
	施蒂里亚Styria 	styria.StyriaDuke
	蒂罗尔Tyrol 	tyrol.TyrolDuke
}

func (k *巴伐利亚BavariaKingdom) DBavaria巴伐利亚() bavaria.BavariaDuke {
	return k.巴伐利亚Bavaria
}
    
func (k *巴伐利亚BavariaKingdom) DNordgau诺德高() nordgau.NordgauDuke {
	return k.诺德高Nordgau
}
    
func (k *巴伐利亚BavariaKingdom) DOsterreich奥地利() osterreich.OsterreichDuke {
	return k.奥地利Osterreich
}
    
func (k *巴伐利亚BavariaKingdom) DSalzburg萨尔茨堡() salzburg.SalzburgDuke {
	return k.萨尔茨堡Salzburg
}
    
func (k *巴伐利亚BavariaKingdom) DStyria施蒂里亚() styria.StyriaDuke {
	return k.施蒂里亚Styria
}
    
func (k *巴伐利亚BavariaKingdom) DTyrol蒂罗尔() tyrol.TyrolDuke {
	return k.蒂罗尔Tyrol
}
    
var KBavaria巴伐利亚 BavariaKingdom = &巴伐利亚BavariaKingdom{}

func init() {
	f := KBavaria巴伐利亚.(*巴伐利亚BavariaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "bavaria",
		TitleName: "巴伐利亚",
		TitleCode: "k_bavaria",
		Dukes:  map[string]feud.Duke{},
	}

	f.巴伐利亚Bavaria = bavaria.DBavaria巴伐利亚
	f.巴伐利亚Bavaria.SetParent(f)
	
	f.诺德高Nordgau = nordgau.DNordgau诺德高
	f.诺德高Nordgau.SetParent(f)
	
	f.奥地利Osterreich = osterreich.DOsterreich奥地利
	f.奥地利Osterreich.SetParent(f)
	
	f.萨尔茨堡Salzburg = salzburg.DSalzburg萨尔茨堡
	f.萨尔茨堡Salzburg.SetParent(f)
	
	f.施蒂里亚Styria = styria.DStyria施蒂里亚
	f.施蒂里亚Styria.SetParent(f)
	
	f.蒂罗尔Tyrol = tyrol.DTyrol蒂罗尔
	f.蒂罗尔Tyrol.SetParent(f)
	
}
