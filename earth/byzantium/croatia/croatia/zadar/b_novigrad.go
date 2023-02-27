package zadar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺维格勒NovigradBarony struct {
	feud.BaseBarony
}

var BNovigrad诺维格勒 feud.Barony = &诺维格勒NovigradBarony{}

func init() {
    f := BNovigrad诺维格勒.(*诺维格勒NovigradBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novigrad",
		TitleName: "诺维格勒",
		TitleCode: "b_novigrad",
	}
}
