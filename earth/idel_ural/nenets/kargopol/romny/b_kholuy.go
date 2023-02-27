package romny

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍卢伊KholuyBarony struct {
	feud.BaseBarony
}

var BKholuy霍卢伊 feud.Barony = &霍卢伊KholuyBarony{}

func init() {
    f := BKholuy霍卢伊.(*霍卢伊KholuyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kholuy",
		TitleName: "霍卢伊",
		TitleCode: "b_kholuy",
	}
}
