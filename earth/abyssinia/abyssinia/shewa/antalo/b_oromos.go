package antalo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥罗莫斯OromosBarony struct {
	feud.BaseBarony
}

var BOromos奥罗莫斯 feud.Barony = &奥罗莫斯OromosBarony{}

func init() {
    f := BOromos奥罗莫斯.(*奥罗莫斯OromosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oromos",
		TitleName: "奥罗莫斯",
		TitleCode: "b_oromos",
	}
}
