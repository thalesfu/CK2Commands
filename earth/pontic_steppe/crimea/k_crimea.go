package crimea

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/cherson"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/crimea"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/crimea/wild_fields"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CrimeaKingdom interface {
    feud.Kingdom
    DCherson赫尔松() 	cherson.ChersonDuke
    DCrimea克里米亚() 	crimea.CrimeaDuke
    DWild_fields艾泰尔克兹() 	wild_fields.Wild_fieldsDuke
}

type 克里米亚CrimeaKingdom struct {
	feud.BaseKingdom
	赫尔松Cherson 	cherson.ChersonDuke
	克里米亚Crimea 	crimea.CrimeaDuke
	艾泰尔克兹Wild_fields 	wild_fields.Wild_fieldsDuke
}

func (k *克里米亚CrimeaKingdom) DCherson赫尔松() cherson.ChersonDuke {
	return k.赫尔松Cherson
}
    
func (k *克里米亚CrimeaKingdom) DCrimea克里米亚() crimea.CrimeaDuke {
	return k.克里米亚Crimea
}
    
func (k *克里米亚CrimeaKingdom) DWild_fields艾泰尔克兹() wild_fields.Wild_fieldsDuke {
	return k.艾泰尔克兹Wild_fields
}
    
var KCrimea克里米亚 CrimeaKingdom = &克里米亚CrimeaKingdom{}

func init() {
	f := KCrimea克里米亚.(*克里米亚CrimeaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "crimea",
		TitleName: "克里米亚",
		TitleCode: "k_crimea",
		Dukes:  map[string]feud.Duke{},
	}

	f.赫尔松Cherson = cherson.DCherson赫尔松
	f.赫尔松Cherson.SetParent(f)
	
	f.克里米亚Crimea = crimea.DCrimea克里米亚
	f.克里米亚Crimea.SetParent(f)
	
	f.艾泰尔克兹Wild_fields = wild_fields.DWild_fields艾泰尔克兹
	f.艾泰尔克兹Wild_fields.SetParent(f)
	
}
