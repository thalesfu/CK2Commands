package warzazat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施旺ShwanBarony struct {
	feud.BaseBarony
}

var BShwan施旺 feud.Barony = &施旺ShwanBarony{}

func init() {
	f := BShwan施旺.(*施旺ShwanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shwan",
		TitleName: "施旺",
		TitleCode: "b_shwan",
	}
}
