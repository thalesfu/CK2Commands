package al_hasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥约恩AloyoonBarony struct {
	feud.BaseBarony
}

var BAloyoon奥约恩 feud.Barony = &奥约恩AloyoonBarony{}

func init() {
    f := BAloyoon奥约恩.(*奥约恩AloyoonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aloyoon",
		TitleName: "奥约恩",
		TitleCode: "b_aloyoon",
	}
}
