package magadha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩内罗ManerBarony struct {
	feud.BaseBarony
}

var BManer摩内罗 feud.Barony = &摩内罗ManerBarony{}

func init() {
	f := BManer摩内罗.(*摩内罗ManerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maner",
		TitleName: "摩内罗",
		TitleCode: "b_maner",
	}
}
