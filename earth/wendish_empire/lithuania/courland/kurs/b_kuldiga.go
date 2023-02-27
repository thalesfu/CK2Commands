package kurs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔迪加KuldigaBarony struct {
	feud.BaseBarony
}

var BKuldiga库尔迪加 feud.Barony = &库尔迪加KuldigaBarony{}

func init() {
    f := BKuldiga库尔迪加.(*库尔迪加KuldigaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuldiga",
		TitleName: "库尔迪加",
		TitleCode: "b_kuldiga",
	}
}
