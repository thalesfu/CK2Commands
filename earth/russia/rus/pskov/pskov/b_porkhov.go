package pskov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波尔霍夫PorkhovBarony struct {
	feud.BaseBarony
}

var BPorkhov波尔霍夫 feud.Barony = &波尔霍夫PorkhovBarony{}

func init() {
	f := BPorkhov波尔霍夫.(*波尔霍夫PorkhovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "porkhov",
		TitleName: "波尔霍夫",
		TitleCode: "b_porkhov",
	}
}
