package burgundy

import (
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/dauphine"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/franche_comte"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/provence"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/savoie"
	"github.com/thalesfu/CK2Commands/earth/france/burgundy/upper_burgundy"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BurgundyKingdom interface {
    feud.Kingdom
    DDauphine多菲内() 	dauphine.DauphineDuke
    DFranche_comte弗朗什孔泰() 	franche_comte.Franche_comteDuke
    DProvence普罗旺斯() 	provence.ProvenceDuke
    DSavoie萨伏依() 	savoie.SavoieDuke
    DUpper_burgundy上勃艮第() 	upper_burgundy.Upper_burgundyDuke
}

type 勃艮第BurgundyKingdom struct {
	feud.BaseKingdom
	多菲内Dauphine 	dauphine.DauphineDuke
	弗朗什孔泰Franche_comte 	franche_comte.Franche_comteDuke
	普罗旺斯Provence 	provence.ProvenceDuke
	萨伏依Savoie 	savoie.SavoieDuke
	上勃艮第Upper_burgundy 	upper_burgundy.Upper_burgundyDuke
}

func (k *勃艮第BurgundyKingdom) DDauphine多菲内() dauphine.DauphineDuke {
	return k.多菲内Dauphine
}
    
func (k *勃艮第BurgundyKingdom) DFranche_comte弗朗什孔泰() franche_comte.Franche_comteDuke {
	return k.弗朗什孔泰Franche_comte
}
    
func (k *勃艮第BurgundyKingdom) DProvence普罗旺斯() provence.ProvenceDuke {
	return k.普罗旺斯Provence
}
    
func (k *勃艮第BurgundyKingdom) DSavoie萨伏依() savoie.SavoieDuke {
	return k.萨伏依Savoie
}
    
func (k *勃艮第BurgundyKingdom) DUpper_burgundy上勃艮第() upper_burgundy.Upper_burgundyDuke {
	return k.上勃艮第Upper_burgundy
}
    
var KBurgundy勃艮第 BurgundyKingdom = &勃艮第BurgundyKingdom{}

func init() {
	f := KBurgundy勃艮第.(*勃艮第BurgundyKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "burgundy",
		TitleName: "勃艮第",
		TitleCode: "k_burgundy",
		Dukes:  map[string]feud.Duke{},
	}

	f.多菲内Dauphine = dauphine.DDauphine多菲内
	f.多菲内Dauphine.SetParent(f)
	
	f.弗朗什孔泰Franche_comte = franche_comte.DFranche_comte弗朗什孔泰
	f.弗朗什孔泰Franche_comte.SetParent(f)
	
	f.普罗旺斯Provence = provence.DProvence普罗旺斯
	f.普罗旺斯Provence.SetParent(f)
	
	f.萨伏依Savoie = savoie.DSavoie萨伏依
	f.萨伏依Savoie.SetParent(f)
	
	f.上勃艮第Upper_burgundy = upper_burgundy.DUpper_burgundy上勃艮第
	f.上勃艮第Upper_burgundy.SetParent(f)
	
}
