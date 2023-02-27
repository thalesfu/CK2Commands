package damascus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔达维勒堡QsarbardawilBarony struct {
	feud.BaseBarony
}

var BQsarbardawil巴尔达维勒堡 feud.Barony = &巴尔达维勒堡QsarbardawilBarony{}

func init() {
    f := BQsarbardawil巴尔达维勒堡.(*巴尔达维勒堡QsarbardawilBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qsarbardawil",
		TitleName: "巴尔达维勒堡",
		TitleCode: "b_qsarbardawil",
	}
}
