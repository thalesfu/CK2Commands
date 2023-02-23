package murom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 穆罗姆MuromBarony struct {
	feud.BaseBarony
}

var BMurom穆罗姆 feud.Barony = &穆罗姆MuromBarony{}

func init() {
	f := BMurom穆罗姆.(*穆罗姆MuromBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "murom",
		TitleName: "穆罗姆",
		TitleCode: "b_murom",
	}
}
