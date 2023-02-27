package brittany

import (
	"github.com/thalesfu/CK2Commands/earth/france/brittany/brittany"
	"github.com/thalesfu/CK2Commands/earth/france/brittany/loire"
	"github.com/thalesfu/CK2Commands/earth/france/brittany/penthievre"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BrittanyKingdom interface {
    feud.Kingdom
    DBrittany布列塔尼() 	brittany.BrittanyDuke
    DLoire上布列塔尼() 	loire.LoireDuke
    DPenthievre庞蒂耶夫尔() 	penthievre.PenthievreDuke
}

type 布列塔尼BrittanyKingdom struct {
	feud.BaseKingdom
	布列塔尼Brittany 	brittany.BrittanyDuke
	上布列塔尼Loire 	loire.LoireDuke
	庞蒂耶夫尔Penthievre 	penthievre.PenthievreDuke
}

func (k *布列塔尼BrittanyKingdom) DBrittany布列塔尼() brittany.BrittanyDuke {
	return k.布列塔尼Brittany
}
    
func (k *布列塔尼BrittanyKingdom) DLoire上布列塔尼() loire.LoireDuke {
	return k.上布列塔尼Loire
}
    
func (k *布列塔尼BrittanyKingdom) DPenthievre庞蒂耶夫尔() penthievre.PenthievreDuke {
	return k.庞蒂耶夫尔Penthievre
}
    
var KBrittany布列塔尼 BrittanyKingdom = &布列塔尼BrittanyKingdom{}

func init() {
	f := KBrittany布列塔尼.(*布列塔尼BrittanyKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "brittany",
		TitleName: "布列塔尼",
		TitleCode: "k_brittany",
		Dukes:  map[string]feud.Duke{},
	}

	f.布列塔尼Brittany = brittany.DBrittany布列塔尼
	f.布列塔尼Brittany.SetParent(f)
	
	f.上布列塔尼Loire = loire.DLoire上布列塔尼
	f.上布列塔尼Loire.SetParent(f)
	
	f.庞蒂耶夫尔Penthievre = penthievre.DPenthievre庞蒂耶夫尔
	f.庞蒂耶夫尔Penthievre.SetParent(f)
	
}
