package sutai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 明安MyanganBarony struct {
	feud.BaseBarony
}

var BMyangan明安 feud.Barony = &明安MyanganBarony{}

func init() {
    f := BMyangan明安.(*明安MyanganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "myangan",
		TitleName: "明安",
		TitleCode: "b_myangan",
	}
}
