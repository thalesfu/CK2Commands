package poznanskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科希强KoscianBarony struct {
	feud.BaseBarony
}

var BKoscian科希强 feud.Barony = &科希强KoscianBarony{}

func init() {
    f := BKoscian科希强.(*科希强KoscianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koscian",
		TitleName: "科希强",
		TitleCode: "b_koscian",
	}
}
