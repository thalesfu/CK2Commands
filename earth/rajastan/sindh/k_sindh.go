package sindh

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/sindh/bhakkar"
	"github.com/thalesfu/CK2Commands/earth/rajastan/sindh/sauvira"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SindhKingdom interface {
    feud.Kingdom
    DBhakkar珀格尔() 	bhakkar.BhakkarDuke
    DSauvira粟毗罗() 	sauvira.SauviraDuke
}

type 信度SindhKingdom struct {
	feud.BaseKingdom
	珀格尔Bhakkar 	bhakkar.BhakkarDuke
	粟毗罗Sauvira 	sauvira.SauviraDuke
}

func (k *信度SindhKingdom) DBhakkar珀格尔() bhakkar.BhakkarDuke {
	return k.珀格尔Bhakkar
}
    
func (k *信度SindhKingdom) DSauvira粟毗罗() sauvira.SauviraDuke {
	return k.粟毗罗Sauvira
}
    
var KSindh信度 SindhKingdom = &信度SindhKingdom{}

func init() {
	f := KSindh信度.(*信度SindhKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "sindh",
		TitleName: "信度",
		TitleCode: "k_sindh",
		Dukes:  map[string]feud.Duke{},
	}

	f.珀格尔Bhakkar = bhakkar.DBhakkar珀格尔
	f.珀格尔Bhakkar.SetParent(f)
	
	f.粟毗罗Sauvira = sauvira.DSauvira粟毗罗
	f.粟毗罗Sauvira.SetParent(f)
	
}
