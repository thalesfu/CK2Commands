package argyll

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔赫ArghBarony struct {
	feud.BaseBarony
}

var BArgh阿尔赫 feud.Barony = &阿尔赫ArghBarony{}

func init() {
    f := BArgh阿尔赫.(*阿尔赫ArghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "argh",
		TitleName: "阿尔赫",
		TitleCode: "b_argh",
	}
}
