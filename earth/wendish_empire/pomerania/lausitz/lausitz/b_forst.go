package lausitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福斯特ForstBarony struct {
	feud.BaseBarony
}

var BForst福斯特 feud.Barony = &福斯特ForstBarony{}

func init() {
    f := BForst福斯特.(*福斯特ForstBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "forst",
		TitleName: "福斯特",
		TitleCode: "b_forst",
	}
}
