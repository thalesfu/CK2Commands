package euboia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊斯泰雅IstiaiaBarony struct {
	feud.BaseBarony
}

var BIstiaia伊斯泰雅 feud.Barony = &伊斯泰雅IstiaiaBarony{}

func init() {
    f := BIstiaia伊斯泰雅.(*伊斯泰雅IstiaiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "istiaia",
		TitleName: "伊斯泰雅",
		TitleCode: "b_istiaia",
	}
}
