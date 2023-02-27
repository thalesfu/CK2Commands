package itil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔加AlgaBarony struct {
	feud.BaseBarony
}

var BAlga阿尔加 feud.Barony = &阿尔加AlgaBarony{}

func init() {
    f := BAlga阿尔加.(*阿尔加AlgaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alga",
		TitleName: "阿尔加",
		TitleCode: "b_alga",
	}
}
