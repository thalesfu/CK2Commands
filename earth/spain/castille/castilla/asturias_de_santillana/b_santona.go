package asturias_de_santillana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑托尼亚SantonaBarony struct {
	feud.BaseBarony
}

var BSantona桑托尼亚 feud.Barony = &桑托尼亚SantonaBarony{}

func init() {
    f := BSantona桑托尼亚.(*桑托尼亚SantonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santona",
		TitleName: "桑托尼亚",
		TitleCode: "b_santona",
	}
}
