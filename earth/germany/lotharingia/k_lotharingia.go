package lotharingia

import (
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/brabant"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/koln"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/lower_lorraine"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/trier"
	"github.com/thalesfu/CK2Commands/earth/germany/lotharingia/upper_lorraine"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LotharingiaKingdom interface {
    feud.Kingdom
    DBrabant布拉班特() 	brabant.BrabantDuke
    DKoln科隆() 	koln.KolnDuke
    DLower_lorraine下洛林() 	lower_lorraine.Lower_lorraineDuke
    DTrier特里尔() 	trier.TrierDuke
    DUpper_lorraine上洛林() 	upper_lorraine.Upper_lorraineDuke
}

type 中法兰克LotharingiaKingdom struct {
	feud.BaseKingdom
	布拉班特Brabant 	brabant.BrabantDuke
	科隆Koln 	koln.KolnDuke
	下洛林Lower_lorraine 	lower_lorraine.Lower_lorraineDuke
	特里尔Trier 	trier.TrierDuke
	上洛林Upper_lorraine 	upper_lorraine.Upper_lorraineDuke
}

func (k *中法兰克LotharingiaKingdom) DBrabant布拉班特() brabant.BrabantDuke {
	return k.布拉班特Brabant
}
    
func (k *中法兰克LotharingiaKingdom) DKoln科隆() koln.KolnDuke {
	return k.科隆Koln
}
    
func (k *中法兰克LotharingiaKingdom) DLower_lorraine下洛林() lower_lorraine.Lower_lorraineDuke {
	return k.下洛林Lower_lorraine
}
    
func (k *中法兰克LotharingiaKingdom) DTrier特里尔() trier.TrierDuke {
	return k.特里尔Trier
}
    
func (k *中法兰克LotharingiaKingdom) DUpper_lorraine上洛林() upper_lorraine.Upper_lorraineDuke {
	return k.上洛林Upper_lorraine
}
    
var KLotharingia中法兰克 LotharingiaKingdom = &中法兰克LotharingiaKingdom{}

func init() {
	f := KLotharingia中法兰克.(*中法兰克LotharingiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "lotharingia",
		TitleName: "中法兰克",
		TitleCode: "k_lotharingia",
		Dukes:  map[string]feud.Duke{},
	}

	f.布拉班特Brabant = brabant.DBrabant布拉班特
	f.布拉班特Brabant.SetParent(f)
	
	f.科隆Koln = koln.DKoln科隆
	f.科隆Koln.SetParent(f)
	
	f.下洛林Lower_lorraine = lower_lorraine.DLower_lorraine下洛林
	f.下洛林Lower_lorraine.SetParent(f)
	
	f.特里尔Trier = trier.DTrier特里尔
	f.特里尔Trier.SetParent(f)
	
	f.上洛林Upper_lorraine = upper_lorraine.DUpper_lorraine上洛林
	f.上洛林Upper_lorraine.SetParent(f)
	
}
