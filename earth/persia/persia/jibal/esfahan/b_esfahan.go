package esfahan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯法罕EsfahanBarony struct {
	feud.BaseBarony
}

var BEsfahan伊斯法罕 feud.Barony = &伊斯法罕EsfahanBarony{}

func init() {
	f := BEsfahan伊斯法罕.(*伊斯法罕EsfahanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "esfahan",
		TitleName: "伊斯法罕",
		TitleCode: "b_esfahan",
	}
}
