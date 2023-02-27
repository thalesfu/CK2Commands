package desmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科克CorkBarony struct {
	feud.BaseBarony
}

var BCork科克 feud.Barony = &科克CorkBarony{}

func init() {
    f := BCork科克.(*科克CorkBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cork",
		TitleName: "科克",
		TitleCode: "b_cork",
	}
}
