package karluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙格阔勒ShagkolBarony struct {
	feud.BaseBarony
}

var BShagkol沙格阔勒 feud.Barony = &沙格阔勒ShagkolBarony{}

func init() {
    f := BShagkol沙格阔勒.(*沙格阔勒ShagkolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shagkol",
		TitleName: "沙格阔勒",
		TitleCode: "b_shagkol",
	}
}
