package vannes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 欧赖AurayBarony struct {
	feud.BaseBarony
}

var BAuray欧赖 feud.Barony = &欧赖AurayBarony{}

func init() {
    f := BAuray欧赖.(*欧赖AurayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "auray",
		TitleName: "欧赖",
		TitleCode: "b_auray",
	}
}
