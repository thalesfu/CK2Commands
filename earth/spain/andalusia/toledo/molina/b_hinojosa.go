package molina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊诺霍萨HinojosaBarony struct {
	feud.BaseBarony
}

var BHinojosa伊诺霍萨 feud.Barony = &伊诺霍萨HinojosaBarony{}

func init() {
    f := BHinojosa伊诺霍萨.(*伊诺霍萨HinojosaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hinojosa",
		TitleName: "伊诺霍萨",
		TitleCode: "b_hinojosa",
	}
}
