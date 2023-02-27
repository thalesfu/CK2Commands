package dorset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科夫CorfeBarony struct {
	feud.BaseBarony
}

var BCorfe科夫 feud.Barony = &科夫CorfeBarony{}

func init() {
    f := BCorfe科夫.(*科夫CorfeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "corfe",
		TitleName: "科夫",
		TitleCode: "b_corfe",
	}
}
