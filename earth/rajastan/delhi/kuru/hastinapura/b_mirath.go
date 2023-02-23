package hastinapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弥罗多MirathBarony struct {
	feud.BaseBarony
}

var BMirath弥罗多 feud.Barony = &弥罗多MirathBarony{}

func init() {
	f := BMirath弥罗多.(*弥罗多MirathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mirath",
		TitleName: "弥罗多",
		TitleCode: "b_mirath",
	}
}
