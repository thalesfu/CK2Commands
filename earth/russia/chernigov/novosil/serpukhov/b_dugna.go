package serpukhov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜格纳DugnaBarony struct {
	feud.BaseBarony
}

var BDugna杜格纳 feud.Barony = &杜格纳DugnaBarony{}

func init() {
	f := BDugna杜格纳.(*杜格纳DugnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dugna",
		TitleName: "杜格纳",
		TitleCode: "b_dugna",
	}
}
