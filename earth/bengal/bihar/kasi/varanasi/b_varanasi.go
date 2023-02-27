package varanasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波罗奈VaranasiBarony struct {
	feud.BaseBarony
}

var BVaranasi波罗奈 feud.Barony = &波罗奈VaranasiBarony{}

func init() {
    f := BVaranasi波罗奈.(*波罗奈VaranasiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varanasi",
		TitleName: "波罗奈",
		TitleCode: "b_varanasi",
	}
}
