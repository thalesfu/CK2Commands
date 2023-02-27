package euphrates

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿米力AmirliBarony struct {
	feud.BaseBarony
}

var BAmirli阿米力 feud.Barony = &阿米力AmirliBarony{}

func init() {
    f := BAmirli阿米力.(*阿米力AmirliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amirli",
		TitleName: "阿米力",
		TitleCode: "b_amirli",
	}
}
