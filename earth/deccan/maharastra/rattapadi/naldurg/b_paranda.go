package naldurg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波烂陀ParandaBarony struct {
	feud.BaseBarony
}

var BParanda波烂陀 feud.Barony = &波烂陀ParandaBarony{}

func init() {
    f := BParanda波烂陀.(*波烂陀ParandaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paranda",
		TitleName: "波烂陀",
		TitleCode: "b_paranda",
	}
}
