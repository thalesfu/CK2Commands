package desmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔不列颠KilbrittainBarony struct {
	feud.BaseBarony
}

var BKilbrittain基尔不列颠 feud.Barony = &基尔不列颠KilbrittainBarony{}

func init() {
    f := BKilbrittain基尔不列颠.(*基尔不列颠KilbrittainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kilbrittain",
		TitleName: "基尔不列颠",
		TitleCode: "b_kilbrittain",
	}
}
