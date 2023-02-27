package asir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨达SadahBarony struct {
	feud.BaseBarony
}

var BSadah萨达 feud.Barony = &萨达SadahBarony{}

func init() {
    f := BSadah萨达.(*萨达SadahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sadah",
		TitleName: "萨达",
		TitleCode: "b_sadah",
	}
}
