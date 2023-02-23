package kunlun

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 邦达BangdagBarony struct {
	feud.BaseBarony
}

var BBangdag邦达 feud.Barony = &邦达BangdagBarony{}

func init() {
	f := BBangdag邦达.(*邦达BangdagBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bangdag",
		TitleName: "邦达",
		TitleCode: "b_bangdag",
	}
}
