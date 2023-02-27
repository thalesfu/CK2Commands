package blekinge

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利斯特ListerBarony struct {
	feud.BaseBarony
}

var BLister利斯特 feud.Barony = &利斯特ListerBarony{}

func init() {
    f := BLister利斯特.(*利斯特ListerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lister",
		TitleName: "利斯特",
		TitleCode: "b_lister",
	}
}
