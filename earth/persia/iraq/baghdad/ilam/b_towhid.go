package ilam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔胡德TowhidBarony struct {
	feud.BaseBarony
}

var BTowhid塔胡德 feud.Barony = &塔胡德TowhidBarony{}

func init() {
	f := BTowhid塔胡德.(*塔胡德TowhidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "towhid",
		TitleName: "塔胡德",
		TitleCode: "b_towhid",
	}
}
