package teluch

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科克森KoksenBarony struct {
	feud.BaseBarony
}

var BKoksen科克森 feud.Barony = &科克森KoksenBarony{}

func init() {
    f := BKoksen科克森.(*科克森KoksenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koksen",
		TitleName: "科克森",
		TitleCode: "b_koksen",
	}
}
