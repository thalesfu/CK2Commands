package euboia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥利奥OreoiBarony struct {
	feud.BaseBarony
}

var BOreoi奥利奥 feud.Barony = &奥利奥OreoiBarony{}

func init() {
    f := BOreoi奥利奥.(*奥利奥OreoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oreoi",
		TitleName: "奥利奥",
		TitleCode: "b_oreoi",
	}
}
