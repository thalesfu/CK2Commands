package pecheneg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨多维SadovyBarony struct {
	feud.BaseBarony
}

var BSadovy萨多维 feud.Barony = &萨多维SadovyBarony{}

func init() {
    f := BSadovy萨多维.(*萨多维SadovyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sadovy",
		TitleName: "萨多维",
		TitleCode: "b_sadovy",
	}
}
