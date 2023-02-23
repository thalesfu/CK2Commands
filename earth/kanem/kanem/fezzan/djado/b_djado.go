package djado

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贾多DjadoBarony struct {
	feud.BaseBarony
}

var BDjado贾多 feud.Barony = &贾多DjadoBarony{}

func init() {
	f := BDjado贾多.(*贾多DjadoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "djado",
		TitleName: "贾多",
		TitleCode: "b_djado",
	}
}
