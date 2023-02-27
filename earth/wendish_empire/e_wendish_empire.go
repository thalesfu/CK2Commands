package wendish_empire

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/bohemia"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/lithuania"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/moravia"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/poland"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/pomerania"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Wendish_empireEmpire interface {
    feud.Empire
    KBohemia波希米亚() 	bohemia.BohemiaKingdom
    KLithuania立陶宛() 	lithuania.LithuaniaKingdom
    KMoravia大摩拉维亚() 	moravia.MoraviaKingdom
    KPoland波兰() 	poland.PolandKingdom
    KPomerania波美拉尼亚() 	pomerania.PomeraniaKingdom
}

type 文德帝国Wendish_empireEmpire struct {
	feud.BaseEmpire
	波希米亚Bohemia 	bohemia.BohemiaKingdom
	立陶宛Lithuania 	lithuania.LithuaniaKingdom
	大摩拉维亚Moravia 	moravia.MoraviaKingdom
	波兰Poland 	poland.PolandKingdom
	波美拉尼亚Pomerania 	pomerania.PomeraniaKingdom
}

func (e *文德帝国Wendish_empireEmpire) KBohemia波希米亚() bohemia.BohemiaKingdom {
	return e.波希米亚Bohemia
}
    
func (e *文德帝国Wendish_empireEmpire) KLithuania立陶宛() lithuania.LithuaniaKingdom {
	return e.立陶宛Lithuania
}
    
func (e *文德帝国Wendish_empireEmpire) KMoravia大摩拉维亚() moravia.MoraviaKingdom {
	return e.大摩拉维亚Moravia
}
    
func (e *文德帝国Wendish_empireEmpire) KPoland波兰() poland.PolandKingdom {
	return e.波兰Poland
}
    
func (e *文德帝国Wendish_empireEmpire) KPomerania波美拉尼亚() pomerania.PomeraniaKingdom {
	return e.波美拉尼亚Pomerania
}
    
var EWendish_empire文德帝国 Wendish_empireEmpire = &文德帝国Wendish_empireEmpire{}

func init() {
	f := EWendish_empire文德帝国.(*文德帝国Wendish_empireEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "wendish_empire",
		TitleName: "文德帝国",
		TitleCode: "e_wendish_empire",
		Kingdoms:  map[string]feud.Kingdom{},
	}
	f.波希米亚Bohemia = bohemia.KBohemia波希米亚
	f.波希米亚Bohemia.SetParent(f)
	f.立陶宛Lithuania = lithuania.KLithuania立陶宛
	f.立陶宛Lithuania.SetParent(f)
	f.大摩拉维亚Moravia = moravia.KMoravia大摩拉维亚
	f.大摩拉维亚Moravia.SetParent(f)
	f.波兰Poland = poland.KPoland波兰
	f.波兰Poland.SetParent(f)
	f.波美拉尼亚Pomerania = pomerania.KPomerania波美拉尼亚
	f.波美拉尼亚Pomerania.SetParent(f)
}
