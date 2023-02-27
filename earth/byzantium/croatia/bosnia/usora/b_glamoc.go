package usora

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉莫奇GlamocBarony struct {
	feud.BaseBarony
}

var BGlamoc格拉莫奇 feud.Barony = &格拉莫奇GlamocBarony{}

func init() {
    f := BGlamoc格拉莫奇.(*格拉莫奇GlamocBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glamoc",
		TitleName: "格拉莫奇",
		TitleCode: "b_glamoc",
	}
}
