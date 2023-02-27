package feher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙杰涅德NagyenyedBarony struct {
	feud.BaseBarony
}

var BNagyenyed瑙杰涅德 feud.Barony = &瑙杰涅德NagyenyedBarony{}

func init() {
    f := BNagyenyed瑙杰涅德.(*瑙杰涅德NagyenyedBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nagyenyed",
		TitleName: "瑙杰涅德",
		TitleCode: "b_nagyenyed",
	}
}
