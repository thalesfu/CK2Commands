package parma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲登扎FidenzaBarony struct {
	feud.BaseBarony
}

var BFidenza菲登扎 feud.Barony = &菲登扎FidenzaBarony{}

func init() {
    f := BFidenza菲登扎.(*菲登扎FidenzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fidenza",
		TitleName: "菲登扎",
		TitleCode: "b_fidenza",
	}
}
