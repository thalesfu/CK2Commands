package tinmallal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿格马特AghmatBarony struct {
	feud.BaseBarony
}

var BAghmat阿格马特 feud.Barony = &阿格马特AghmatBarony{}

func init() {
    f := BAghmat阿格马特.(*阿格马特AghmatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aghmat",
		TitleName: "阿格马特",
		TitleCode: "b_aghmat",
	}
}
