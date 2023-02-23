package gurgi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨多那SatnaBarony struct {
	feud.BaseBarony
}

var BSatna萨多那 feud.Barony = &萨多那SatnaBarony{}

func init() {
	f := BSatna萨多那.(*萨多那SatnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "satna",
		TitleName: "萨多那",
		TitleCode: "b_satna",
	}
}
