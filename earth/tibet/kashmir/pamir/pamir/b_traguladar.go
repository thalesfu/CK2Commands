package pamir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉古拉达尔TraguladarBarony struct {
	feud.BaseBarony
}

var BTraguladar特拉古拉达尔 feud.Barony = &特拉古拉达尔TraguladarBarony{}

func init() {
	f := BTraguladar特拉古拉达尔.(*特拉古拉达尔TraguladarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "traguladar",
		TitleName: "特拉古拉达尔",
		TitleCode: "b_traguladar",
	}
}
