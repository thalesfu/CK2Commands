package tartaria

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman"
	"github.com/thalesfu/CK2Commands/earth/tartaria/mongolia"
	"github.com/thalesfu/CK2Commands/earth/tartaria/sibir"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TartariaEmpire interface {
    feud.Empire
    KCuman库曼() 	cuman.CumanKingdom
    KMongolia蒙古() 	mongolia.MongoliaKingdom
    KSibir失必儿() 	sibir.SibirKingdom
}

type 鞑靼TartariaEmpire struct {
	feud.BaseEmpire
	库曼Cuman 	cuman.CumanKingdom
	蒙古Mongolia 	mongolia.MongoliaKingdom
	失必儿Sibir 	sibir.SibirKingdom
}

func (e *鞑靼TartariaEmpire) KCuman库曼() cuman.CumanKingdom {
	return e.库曼Cuman
}
    
func (e *鞑靼TartariaEmpire) KMongolia蒙古() mongolia.MongoliaKingdom {
	return e.蒙古Mongolia
}
    
func (e *鞑靼TartariaEmpire) KSibir失必儿() sibir.SibirKingdom {
	return e.失必儿Sibir
}
    
var ETartaria鞑靼 TartariaEmpire = &鞑靼TartariaEmpire{}

func init() {
	f := ETartaria鞑靼.(*鞑靼TartariaEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "tartaria",
		TitleName: "鞑靼",
		TitleCode: "e_tartaria",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.库曼Cuman = cuman.KCuman库曼
	f.库曼Cuman.SetParent(f)
	f.蒙古Mongolia = mongolia.KMongolia蒙古
	f.蒙古Mongolia.SetParent(f)
	f.失必儿Sibir = sibir.KSibir失必儿
	f.失必儿Sibir.SetParent(f)
}
