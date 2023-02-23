package egitanea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊达尼亚IdanhaBarony struct {
	feud.BaseBarony
}

var BIdanha伊达尼亚 feud.Barony = &伊达尼亚IdanhaBarony{}

func init() {
	f := BIdanha伊达尼亚.(*伊达尼亚IdanhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "idanha",
		TitleName: "伊达尼亚",
		TitleCode: "b_idanha",
	}
}
