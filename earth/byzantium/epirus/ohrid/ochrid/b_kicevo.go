package ochrid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基切沃KicevoBarony struct {
	feud.BaseBarony
}

var BKicevo基切沃 feud.Barony = &基切沃KicevoBarony{}

func init() {
	f := BKicevo基切沃.(*基切沃KicevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kicevo",
		TitleName: "基切沃",
		TitleCode: "b_kicevo",
	}
}
