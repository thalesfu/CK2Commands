package corfu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科孚CorfuBarony struct {
	feud.BaseBarony
}

var BCorfu科孚 feud.Barony = &科孚CorfuBarony{}

func init() {
    f := BCorfu科孚.(*科孚CorfuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "corfu",
		TitleName: "科孚",
		TitleCode: "b_corfu",
	}
}
