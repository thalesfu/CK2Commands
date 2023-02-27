package julich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖尔登GeldernBarony struct {
	feud.BaseBarony
}

var BGeldern盖尔登 feud.Barony = &盖尔登GeldernBarony{}

func init() {
    f := BGeldern盖尔登.(*盖尔登GeldernBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "geldern",
		TitleName: "盖尔登",
		TitleCode: "b_geldern",
	}
}
