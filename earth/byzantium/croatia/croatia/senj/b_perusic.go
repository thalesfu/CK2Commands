package senj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩鲁希奇PerusicBarony struct {
	feud.BaseBarony
}

var BPerusic佩鲁希奇 feud.Barony = &佩鲁希奇PerusicBarony{}

func init() {
    f := BPerusic佩鲁希奇.(*佩鲁希奇PerusicBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perusic",
		TitleName: "佩鲁希奇",
		TitleCode: "b_perusic",
	}
}
