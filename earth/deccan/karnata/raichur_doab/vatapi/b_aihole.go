package vatapi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿赫勒AiholeBarony struct {
	feud.BaseBarony
}

var BAihole阿赫勒 feud.Barony = &阿赫勒AiholeBarony{}

func init() {
    f := BAihole阿赫勒.(*阿赫勒AiholeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aihole",
		TitleName: "阿赫勒",
		TitleCode: "b_aihole",
	}
}
