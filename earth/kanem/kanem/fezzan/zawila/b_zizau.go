package zawila

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰扎乌ZizauBarony struct {
	feud.BaseBarony
}

var BZizau杰扎乌 feud.Barony = &杰扎乌ZizauBarony{}

func init() {
	f := BZizau杰扎乌.(*杰扎乌ZizauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zizau",
		TitleName: "杰扎乌",
		TitleCode: "b_zizau",
	}
}
