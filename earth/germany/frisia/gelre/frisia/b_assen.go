package frisia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿森AssenBarony struct {
	feud.BaseBarony
}

var BAssen阿森 feud.Barony = &阿森AssenBarony{}

func init() {
	f := BAssen阿森.(*阿森AssenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "assen",
		TitleName: "阿森",
		TitleCode: "b_assen",
	}
}
