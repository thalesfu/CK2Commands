package zawila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙瓦SciauaBarony struct {
	feud.BaseBarony
}

var BSciaua沙瓦 feud.Barony = &沙瓦SciauaBarony{}

func init() {
    f := BSciaua沙瓦.(*沙瓦SciauaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sciaua",
		TitleName: "沙瓦",
		TitleCode: "b_sciaua",
	}
}
