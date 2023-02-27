package bohemia

import (
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/bohemia/bohemia"
	"github.com/thalesfu/CK2Commands/earth/wendish_empire/bohemia/moravia"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BohemiaKingdom interface {
    feud.Kingdom
    DBohemia波希米亚() 	bohemia.BohemiaDuke
    DMoravia摩拉维亚() 	moravia.MoraviaDuke
}

type 波希米亚BohemiaKingdom struct {
	feud.BaseKingdom
	波希米亚Bohemia 	bohemia.BohemiaDuke
	摩拉维亚Moravia 	moravia.MoraviaDuke
}

func (k *波希米亚BohemiaKingdom) DBohemia波希米亚() bohemia.BohemiaDuke {
	return k.波希米亚Bohemia
}
    
func (k *波希米亚BohemiaKingdom) DMoravia摩拉维亚() moravia.MoraviaDuke {
	return k.摩拉维亚Moravia
}
    
var KBohemia波希米亚 BohemiaKingdom = &波希米亚BohemiaKingdom{}

func init() {
	f := KBohemia波希米亚.(*波希米亚BohemiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "bohemia",
		TitleName: "波希米亚",
		TitleCode: "k_bohemia",
		Dukes:  map[string]feud.Duke{},
	}

	f.波希米亚Bohemia = bohemia.DBohemia波希米亚
	f.波希米亚Bohemia.SetParent(f)
	
	f.摩拉维亚Moravia = moravia.DMoravia摩拉维亚
	f.摩拉维亚Moravia.SetParent(f)
	
}
