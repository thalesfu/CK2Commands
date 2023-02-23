package sijilmasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊尔富德ErtoudBarony struct {
	feud.BaseBarony
}

var BErtoud伊尔富德 feud.Barony = &伊尔富德ErtoudBarony{}

func init() {
	f := BErtoud伊尔富德.(*伊尔富德ErtoudBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ertoud",
		TitleName: "伊尔富德",
		TitleCode: "b_ertoud",
	}
}
