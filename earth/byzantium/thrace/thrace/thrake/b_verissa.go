package thrake

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维里萨VerissaBarony struct {
	feud.BaseBarony
}

var BVerissa维里萨 feud.Barony = &维里萨VerissaBarony{}

func init() {
	f := BVerissa维里萨.(*维里萨VerissaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "verissa",
		TitleName: "维里萨",
		TitleCode: "b_verissa",
	}
}
