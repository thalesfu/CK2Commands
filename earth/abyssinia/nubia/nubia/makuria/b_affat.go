package makuria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿夫法特AffatBarony struct {
	feud.BaseBarony
}

var BAffat阿夫法特 feud.Barony = &阿夫法特AffatBarony{}

func init() {
	f := BAffat阿夫法特.(*阿夫法特AffatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "affat",
		TitleName: "阿夫法特",
		TitleCode: "b_affat",
	}
}
