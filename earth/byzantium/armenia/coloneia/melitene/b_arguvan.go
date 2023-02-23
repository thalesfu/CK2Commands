package melitene

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔古万ArguvanBarony struct {
	feud.BaseBarony
}

var BArguvan阿尔古万 feud.Barony = &阿尔古万ArguvanBarony{}

func init() {
	f := BArguvan阿尔古万.(*阿尔古万ArguvanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arguvan",
		TitleName: "阿尔古万",
		TitleCode: "b_arguvan",
	}
}
