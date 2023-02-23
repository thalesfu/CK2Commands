package franconia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FranconiaKingdom interface {
	feud.Kingdom
}

type 法兰克尼亚FranconiaKingdom struct {
	feud.BaseKingdom
}

var KFranconia法兰克尼亚 FranconiaKingdom = &法兰克尼亚FranconiaKingdom{}

func init() {
	f := KFranconia法兰克尼亚.(*法兰克尼亚FranconiaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "franconia",
		TitleName: "法兰克尼亚",
		TitleCode: "k_franconia",
		Dukes:     map[string]feud.Duke{},
	}

}
