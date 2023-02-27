package tudgha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西迪比迪Sidi_bidiBarony struct {
	feud.BaseBarony
}

var BSidi_bidi西迪比迪 feud.Barony = &西迪比迪Sidi_bidiBarony{}

func init() {
    f := BSidi_bidi西迪比迪.(*西迪比迪Sidi_bidiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidi_bidi",
		TitleName: "西迪比迪",
		TitleCode: "b_sidi_bidi",
	}
}
