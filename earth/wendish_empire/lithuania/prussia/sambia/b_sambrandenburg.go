package sambia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 勃兰登堡SambrandenburgBarony struct {
	feud.BaseBarony
}

var BSambrandenburg勃兰登堡 feud.Barony = &勃兰登堡SambrandenburgBarony{}

func init() {
    f := BSambrandenburg勃兰登堡.(*勃兰登堡SambrandenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sambrandenburg",
		TitleName: "勃兰登堡",
		TitleCode: "b_sambrandenburg",
	}
}
