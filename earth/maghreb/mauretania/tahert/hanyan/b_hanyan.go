package hanyan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汗燕HanyanBarony struct {
	feud.BaseBarony
}

var BHanyan汗燕 feud.Barony = &汗燕HanyanBarony{}

func init() {
    f := BHanyan汗燕.(*汗燕HanyanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hanyan",
		TitleName: "汗燕",
		TitleCode: "b_hanyan",
	}
}
