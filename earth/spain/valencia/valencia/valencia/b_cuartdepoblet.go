package valencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 夸尔特德波布莱特CuartdepobletBarony struct {
	feud.BaseBarony
}

var BCuartdepoblet夸尔特德波布莱特 feud.Barony = &夸尔特德波布莱特CuartdepobletBarony{}

func init() {
    f := BCuartdepoblet夸尔特德波布莱特.(*夸尔特德波布莱特CuartdepobletBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cuartdepoblet",
		TitleName: "夸尔特德波布莱特",
		TitleCode: "b_cuartdepoblet",
	}
}
