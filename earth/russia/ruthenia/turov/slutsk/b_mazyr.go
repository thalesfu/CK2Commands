package slutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫济里MazyrBarony struct {
	feud.BaseBarony
}

var BMazyr莫济里 feud.Barony = &莫济里MazyrBarony{}

func init() {
    f := BMazyr莫济里.(*莫济里MazyrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mazyr",
		TitleName: "莫济里",
		TitleCode: "b_mazyr",
	}
}
