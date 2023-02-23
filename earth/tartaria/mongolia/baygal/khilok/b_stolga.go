package khilok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯托尔加StolgaBarony struct {
	feud.BaseBarony
}

var BStolga斯托尔加 feud.Barony = &斯托尔加StolgaBarony{}

func init() {
	f := BStolga斯托尔加.(*斯托尔加StolgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stolga",
		TitleName: "斯托尔加",
		TitleCode: "b_stolga",
	}
}
