package rostov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利维LvyBarony struct {
	feud.BaseBarony
}

var BLvy利维 feud.Barony = &利维LvyBarony{}

func init() {
	f := BLvy利维.(*利维LvyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lvy",
		TitleName: "利维",
		TitleCode: "b_lvy",
	}
}
