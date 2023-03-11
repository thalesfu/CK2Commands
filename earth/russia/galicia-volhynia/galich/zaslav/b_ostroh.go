package zaslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥斯特罗赫OstrohBarony struct {
	feud.BaseBarony
}

var BOstroh奥斯特罗赫 feud.Barony = &奥斯特罗赫OstrohBarony{}

func init() {
    f := BOstroh奥斯特罗赫.(*奥斯特罗赫OstrohBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ostroh",
		TitleName: "奥斯特罗赫",
		TitleCode: "b_ostroh",
	}
}
