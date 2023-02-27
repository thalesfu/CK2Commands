package otrar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比尔利克BirlikBarony struct {
	feud.BaseBarony
}

var BBirlik比尔利克 feud.Barony = &比尔利克BirlikBarony{}

func init() {
    f := BBirlik比尔利克.(*比尔利克BirlikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "birlik",
		TitleName: "比尔利克",
		TitleCode: "b_birlik",
	}
}
