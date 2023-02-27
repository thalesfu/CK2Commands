package castelo_branco

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓜达GuardaBarony struct {
	feud.BaseBarony
}

var BGuarda瓜达 feud.Barony = &瓜达GuardaBarony{}

func init() {
    f := BGuarda瓜达.(*瓜达GuardaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guarda",
		TitleName: "瓜达",
		TitleCode: "b_guarda",
	}
}
