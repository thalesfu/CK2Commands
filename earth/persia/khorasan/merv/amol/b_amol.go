package amol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿莫勒AmolBarony struct {
	feud.BaseBarony
}

var BAmol阿莫勒 feud.Barony = &阿莫勒AmolBarony{}

func init() {
	f := BAmol阿莫勒.(*阿莫勒AmolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amol",
		TitleName: "阿莫勒",
		TitleCode: "b_amol",
	}
}
