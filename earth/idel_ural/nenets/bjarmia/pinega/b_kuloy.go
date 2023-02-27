package pinega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库洛伊KuloyBarony struct {
	feud.BaseBarony
}

var BKuloy库洛伊 feud.Barony = &库洛伊KuloyBarony{}

func init() {
    f := BKuloy库洛伊.(*库洛伊KuloyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuloy",
		TitleName: "库洛伊",
		TitleCode: "b_kuloy",
	}
}
