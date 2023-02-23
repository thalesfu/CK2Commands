package muluja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿德加EladerghaBarony struct {
	feud.BaseBarony
}

var BEladergha阿德加 feud.Barony = &阿德加EladerghaBarony{}

func init() {
	f := BEladergha阿德加.(*阿德加EladerghaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eladergha",
		TitleName: "阿德加",
		TitleCode: "b_eladergha",
	}
}
