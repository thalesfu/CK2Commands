package vas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃什堡VasvarBarony struct {
	feud.BaseBarony
}

var BVasvar沃什堡 feud.Barony = &沃什堡VasvarBarony{}

func init() {
	f := BVasvar沃什堡.(*沃什堡VasvarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vasvar",
		TitleName: "沃什堡",
		TitleCode: "b_vasvar",
	}
}
