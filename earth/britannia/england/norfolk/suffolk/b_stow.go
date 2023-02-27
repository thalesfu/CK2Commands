package suffolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯托StowBarony struct {
	feud.BaseBarony
}

var BStow斯托 feud.Barony = &斯托StowBarony{}

func init() {
    f := BStow斯托.(*斯托StowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stow",
		TitleName: "斯托",
		TitleCode: "b_stow",
	}
}
