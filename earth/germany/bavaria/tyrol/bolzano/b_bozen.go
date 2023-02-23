package bolzano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔扎诺BozenBarony struct {
	feud.BaseBarony
}

var BBozen博尔扎诺 feud.Barony = &博尔扎诺BozenBarony{}

func init() {
	f := BBozen博尔扎诺.(*博尔扎诺BozenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bozen",
		TitleName: "博尔扎诺",
		TitleCode: "b_bozen",
	}
}
