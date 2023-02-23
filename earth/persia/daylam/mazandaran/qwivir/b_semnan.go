package qwivir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞姆南SemnanBarony struct {
	feud.BaseBarony
}

var BSemnan塞姆南 feud.Barony = &塞姆南SemnanBarony{}

func init() {
	f := BSemnan塞姆南.(*塞姆南SemnanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "semnan",
		TitleName: "塞姆南",
		TitleCode: "b_semnan",
	}
}
