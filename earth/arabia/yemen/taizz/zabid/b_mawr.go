package zabid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 毛尔MawrBarony struct {
	feud.BaseBarony
}

var BMawr毛尔 feud.Barony = &毛尔MawrBarony{}

func init() {
	f := BMawr毛尔.(*毛尔MawrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mawr",
		TitleName: "毛尔",
		TitleCode: "b_mawr",
	}
}
