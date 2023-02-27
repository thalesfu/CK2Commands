package lausitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 齐陶ZittauBarony struct {
	feud.BaseBarony
}

var BZittau齐陶 feud.Barony = &齐陶ZittauBarony{}

func init() {
    f := BZittau齐陶.(*齐陶ZittauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zittau",
		TitleName: "齐陶",
		TitleCode: "b_zittau",
	}
}
