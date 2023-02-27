package theodosia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗南FunanBarony struct {
	feud.BaseBarony
}

var BFunan弗南 feud.Barony = &弗南FunanBarony{}

func init() {
    f := BFunan弗南.(*弗南FunanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "funan",
		TitleName: "弗南",
		TitleCode: "b_funan",
	}
}
