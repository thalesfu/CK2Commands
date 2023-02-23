package soso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞贝泰SebeteBarony struct {
	feud.BaseBarony
}

var BSebete塞贝泰 feud.Barony = &塞贝泰SebeteBarony{}

func init() {
	f := BSebete塞贝泰.(*塞贝泰SebeteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sebete",
		TitleName: "塞贝泰",
		TitleCode: "b_sebete",
	}
}
