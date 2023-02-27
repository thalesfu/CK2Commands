package leon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨尔达尼亚SaldanaBarony struct {
	feud.BaseBarony
}

var BSaldana萨尔达尼亚 feud.Barony = &萨尔达尼亚SaldanaBarony{}

func init() {
    f := BSaldana萨尔达尼亚.(*萨尔达尼亚SaldanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saldana",
		TitleName: "萨尔达尼亚",
		TitleCode: "b_saldana",
	}
}
