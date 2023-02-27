package gondwana

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/gondwana/dahala"
	"github.com/thalesfu/CK2Commands/earth/bengal/gondwana/ratanpur"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GondwanaKingdom interface {
    feud.Kingdom
    DDahala陀诃罗() 	dahala.DahalaDuke
    DRatanpur剌怛那补罗() 	ratanpur.RatanpurDuke
}

type 共陀婆那GondwanaKingdom struct {
	feud.BaseKingdom
	陀诃罗Dahala 	dahala.DahalaDuke
	剌怛那补罗Ratanpur 	ratanpur.RatanpurDuke
}

func (k *共陀婆那GondwanaKingdom) DDahala陀诃罗() dahala.DahalaDuke {
	return k.陀诃罗Dahala
}
    
func (k *共陀婆那GondwanaKingdom) DRatanpur剌怛那补罗() ratanpur.RatanpurDuke {
	return k.剌怛那补罗Ratanpur
}
    
var KGondwana共陀婆那 GondwanaKingdom = &共陀婆那GondwanaKingdom{}

func init() {
	f := KGondwana共陀婆那.(*共陀婆那GondwanaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "gondwana",
		TitleName: "共陀婆那",
		TitleCode: "k_gondwana",
		Dukes:  map[string]feud.Duke{},
	}

	f.陀诃罗Dahala = dahala.DDahala陀诃罗
	f.陀诃罗Dahala.SetParent(f)
	
	f.剌怛那补罗Ratanpur = ratanpur.DRatanpur剌怛那补罗
	f.剌怛那补罗Ratanpur.SetParent(f)
	
}
