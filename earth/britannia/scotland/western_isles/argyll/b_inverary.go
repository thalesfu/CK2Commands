package argyll

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因弗拉里InveraryBarony struct {
	feud.BaseBarony
}

var BInverary因弗拉里 feud.Barony = &因弗拉里InveraryBarony{}

func init() {
    f := BInverary因弗拉里.(*因弗拉里InveraryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "inverary",
		TitleName: "因弗拉里",
		TitleCode: "b_inverary",
	}
}
