package udine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔梅佐TolmezzoBarony struct {
	feud.BaseBarony
}

var BTolmezzo托尔梅佐 feud.Barony = &托尔梅佐TolmezzoBarony{}

func init() {
    f := BTolmezzo托尔梅佐.(*托尔梅佐TolmezzoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tolmezzo",
		TitleName: "托尔梅佐",
		TitleCode: "b_tolmezzo",
	}
}
