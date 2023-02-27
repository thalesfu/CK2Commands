package khozistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍韦伊泽HoveizehBarony struct {
	feud.BaseBarony
}

var BHoveizeh霍韦伊泽 feud.Barony = &霍韦伊泽HoveizehBarony{}

func init() {
    f := BHoveizeh霍韦伊泽.(*霍韦伊泽HoveizehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hoveizeh",
		TitleName: "霍韦伊泽",
		TitleCode: "b_hoveizeh",
	}
}
