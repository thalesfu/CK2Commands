package kolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库马伊KumayBarony struct {
	feud.BaseBarony
}

var BKumay库马伊 feud.Barony = &库马伊KumayBarony{}

func init() {
    f := BKumay库马伊.(*库马伊KumayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kumay",
		TitleName: "库马伊",
		TitleCode: "b_kumay",
	}
}
