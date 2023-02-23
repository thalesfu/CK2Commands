package candar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CandarKingdom interface {
	feud.Kingdom
}

type 钱达尔CandarKingdom struct {
	feud.BaseKingdom
}

var KCandar钱达尔 CandarKingdom = &钱达尔CandarKingdom{}

func init() {
	f := KCandar钱达尔.(*钱达尔CandarKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "candar",
		TitleName: "钱达尔",
		TitleCode: "k_candar",
		Dukes:     map[string]feud.Duke{},
	}

}
