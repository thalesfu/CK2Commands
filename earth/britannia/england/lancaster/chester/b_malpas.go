package chester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫尔珀斯MalpasBarony struct {
	feud.BaseBarony
}

var BMalpas莫尔珀斯 feud.Barony = &莫尔珀斯MalpasBarony{}

func init() {
	f := BMalpas莫尔珀斯.(*莫尔珀斯MalpasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "malpas",
		TitleName: "莫尔珀斯",
		TitleCode: "b_malpas",
	}
}
