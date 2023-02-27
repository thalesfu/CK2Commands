package yumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙岗ShagangBarony struct {
	feud.BaseBarony
}

var BShagang沙岗 feud.Barony = &沙岗ShagangBarony{}

func init() {
    f := BShagang沙岗.(*沙岗ShagangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shagang",
		TitleName: "沙岗",
		TitleCode: "b_shagang",
	}
}
