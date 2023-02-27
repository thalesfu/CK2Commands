package makran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼恩ManBarony struct {
	feud.BaseBarony
}

var BMan曼恩 feud.Barony = &曼恩ManBarony{}

func init() {
    f := BMan曼恩.(*曼恩ManBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "man",
		TitleName: "曼恩",
		TitleCode: "b_man",
	}
}
