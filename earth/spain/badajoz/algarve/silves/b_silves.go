package silves

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡尔维什SilvesBarony struct {
	feud.BaseBarony
}

var BSilves锡尔维什 feud.Barony = &锡尔维什SilvesBarony{}

func init() {
    f := BSilves锡尔维什.(*锡尔维什SilvesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "silves",
		TitleName: "锡尔维什",
		TitleCode: "b_silves",
	}
}
