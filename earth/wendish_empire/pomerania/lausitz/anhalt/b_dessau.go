package anhalt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德绍DessauBarony struct {
	feud.BaseBarony
}

var BDessau德绍 feud.Barony = &德绍DessauBarony{}

func init() {
    f := BDessau德绍.(*德绍DessauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dessau",
		TitleName: "德绍",
		TitleCode: "b_dessau",
	}
}
