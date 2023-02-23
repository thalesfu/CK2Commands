package parnakheta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦兰阇KaranjaBarony struct {
	feud.BaseBarony
}

var BKaranja迦兰阇 feud.Barony = &迦兰阇KaranjaBarony{}

func init() {
	f := BKaranja迦兰阇.(*迦兰阇KaranjaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karanja",
		TitleName: "迦兰阇",
		TitleCode: "b_karanja",
	}
}
