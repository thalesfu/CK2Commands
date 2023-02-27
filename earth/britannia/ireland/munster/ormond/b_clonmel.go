package ormond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克朗梅尔ClonmelBarony struct {
	feud.BaseBarony
}

var BClonmel克朗梅尔 feud.Barony = &克朗梅尔ClonmelBarony{}

func init() {
    f := BClonmel克朗梅尔.(*克朗梅尔ClonmelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clonmel",
		TitleName: "克朗梅尔",
		TitleCode: "b_clonmel",
	}
}
