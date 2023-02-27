package tell_atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞提夫SetifBarony struct {
	feud.BaseBarony
}

var BSetif塞提夫 feud.Barony = &塞提夫SetifBarony{}

func init() {
    f := BSetif塞提夫.(*塞提夫SetifBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "setif",
		TitleName: "塞提夫",
		TitleCode: "b_setif",
	}
}
