package arques

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔克ArquesBarony struct {
	feud.BaseBarony
}

var BArques阿尔克 feud.Barony = &阿尔克ArquesBarony{}

func init() {
    f := BArques阿尔克.(*阿尔克ArquesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arques",
		TitleName: "阿尔克",
		TitleCode: "b_arques",
	}
}
