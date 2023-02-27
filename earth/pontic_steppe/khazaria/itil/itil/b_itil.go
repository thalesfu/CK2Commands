package itil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿得ItilBarony struct {
	feud.BaseBarony
}

var BItil阿得 feud.Barony = &阿得ItilBarony{}

func init() {
    f := BItil阿得.(*阿得ItilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "itil",
		TitleName: "阿得",
		TitleCode: "b_itil",
	}
}
