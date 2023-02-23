package attaleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯莱治SlegeBarony struct {
	feud.BaseBarony
}

var BSlege斯莱治 feud.Barony = &斯莱治SlegeBarony{}

func init() {
	f := BSlege斯莱治.(*斯莱治SlegeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "slege",
		TitleName: "斯莱治",
		TitleCode: "b_slege",
	}
}
