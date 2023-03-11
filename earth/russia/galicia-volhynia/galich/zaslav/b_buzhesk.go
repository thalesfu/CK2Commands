package zaslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布热斯克BuzheskBarony struct {
	feud.BaseBarony
}

var BBuzhesk布热斯克 feud.Barony = &布热斯克BuzheskBarony{}

func init() {
    f := BBuzhesk布热斯克.(*布热斯克BuzheskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buzhesk",
		TitleName: "布热斯克",
		TitleCode: "b_buzhesk",
	}
}
