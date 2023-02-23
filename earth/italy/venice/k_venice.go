package venice

import (
	"github.com/thalesfu/CK2Commands/earth/italy/venice/venice"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VeniceKingdom interface {
	feud.Kingdom
	DVenice威尼斯() venice.VeniceDuke
}

type 威尼斯VeniceKingdom struct {
	feud.BaseKingdom
	威尼斯Venice venice.VeniceDuke
}

func (k *威尼斯VeniceKingdom) DVenice威尼斯() venice.VeniceDuke {
	return k.威尼斯Venice
}

var KVenice威尼斯 VeniceKingdom = &威尼斯VeniceKingdom{}

func init() {
	f := KVenice威尼斯.(*威尼斯VeniceKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "venice",
		TitleName: "威尼斯",
		TitleCode: "k_venice",
		Dukes:     map[string]feud.Duke{},
	}

	f.威尼斯Venice = venice.DVenice威尼斯
	f.威尼斯Venice.SetParent(f)

}
