package nagadipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿耆阿AgiaBarony struct {
	feud.BaseBarony
}

var BAgia阿耆阿 feud.Barony = &阿耆阿AgiaBarony{}

func init() {
    f := BAgia阿耆阿.(*阿耆阿AgiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agia",
		TitleName: "阿耆阿",
		TitleCode: "b_agia",
	}
}
