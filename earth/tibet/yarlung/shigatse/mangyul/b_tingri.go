package mangyul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 定日TingriBarony struct {
	feud.BaseBarony
}

var BTingri定日 feud.Barony = &定日TingriBarony{}

func init() {
	f := BTingri定日.(*定日TingriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tingri",
		TitleName: "定日",
		TitleCode: "b_tingri",
	}
}
