package lyubech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩列皮斯PerepysBarony struct {
	feud.BaseBarony
}

var BPerepys佩列皮斯 feud.Barony = &佩列皮斯PerepysBarony{}

func init() {
    f := BPerepys佩列皮斯.(*佩列皮斯PerepysBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perepys",
		TitleName: "佩列皮斯",
		TitleCode: "b_perepys",
	}
}
