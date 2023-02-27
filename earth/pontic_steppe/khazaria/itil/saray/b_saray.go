package saray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨赖SarayBarony struct {
	feud.BaseBarony
}

var BSaray萨赖 feud.Barony = &萨赖SarayBarony{}

func init() {
    f := BSaray萨赖.(*萨赖SarayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saray",
		TitleName: "萨赖",
		TitleCode: "b_saray",
	}
}
