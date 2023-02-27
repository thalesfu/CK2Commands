package perfeddwlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗林特FlintBarony struct {
	feud.BaseBarony
}

var BFlint弗林特 feud.Barony = &弗林特FlintBarony{}

func init() {
    f := BFlint弗林特.(*弗林特FlintBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "flint",
		TitleName: "弗林特",
		TitleCode: "b_flint",
	}
}
