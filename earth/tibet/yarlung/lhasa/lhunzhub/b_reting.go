package lhunzhub

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 热振RetingBarony struct {
	feud.BaseBarony
}

var BReting热振 feud.Barony = &热振RetingBarony{}

func init() {
	f := BReting热振.(*热振RetingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reting",
		TitleName: "热振",
		TitleCode: "b_reting",
	}
}
