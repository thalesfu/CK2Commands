package joensuu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃诺EnoBarony struct {
	feud.BaseBarony
}

var BEno埃诺 feud.Barony = &埃诺EnoBarony{}

func init() {
	f := BEno埃诺.(*埃诺EnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eno",
		TitleName: "埃诺",
		TitleCode: "b_eno",
	}
}
