package freistadt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 采尔ZellBarony struct {
	feud.BaseBarony
}

var BZell采尔 feud.Barony = &采尔ZellBarony{}

func init() {
    f := BZell采尔.(*采尔ZellBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zell",
		TitleName: "采尔",
		TitleCode: "b_zell",
	}
}
