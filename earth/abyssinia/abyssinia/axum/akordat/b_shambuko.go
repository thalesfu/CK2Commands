package akordat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙姆布克ShambukoBarony struct {
	feud.BaseBarony
}

var BShambuko沙姆布克 feud.Barony = &沙姆布克ShambukoBarony{}

func init() {
    f := BShambuko沙姆布克.(*沙姆布克ShambukoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shambuko",
		TitleName: "沙姆布克",
		TitleCode: "b_shambuko",
	}
}
