package moskva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫斯科MoskvaBarony struct {
	feud.BaseBarony
}

var BMoskva莫斯科 feud.Barony = &莫斯科MoskvaBarony{}

func init() {
	f := BMoskva莫斯科.(*莫斯科MoskvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moskva",
		TitleName: "莫斯科",
		TitleCode: "b_moskva",
	}
}
