package maymana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔玛勒AlmarBarony struct {
	feud.BaseBarony
}

var BAlmar阿尔玛勒 feud.Barony = &阿尔玛勒AlmarBarony{}

func init() {
    f := BAlmar阿尔玛勒.(*阿尔玛勒AlmarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almar",
		TitleName: "阿尔玛勒",
		TitleCode: "b_almar",
	}
}
