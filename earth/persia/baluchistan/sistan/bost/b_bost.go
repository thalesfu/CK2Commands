package bost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博斯特BostBarony struct {
	feud.BaseBarony
}

var BBost博斯特 feud.Barony = &博斯特BostBarony{}

func init() {
    f := BBost博斯特.(*博斯特BostBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bost",
		TitleName: "博斯特",
		TitleCode: "b_bost",
	}
}
