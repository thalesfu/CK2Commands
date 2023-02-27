package sirjan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙赫尔巴巴克ShahrbabakBarony struct {
	feud.BaseBarony
}

var BShahrbabak沙赫尔巴巴克 feud.Barony = &沙赫尔巴巴克ShahrbabakBarony{}

func init() {
    f := BShahrbabak沙赫尔巴巴克.(*沙赫尔巴巴克ShahrbabakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shahrbabak",
		TitleName: "沙赫尔巴巴克",
		TitleCode: "b_shahrbabak",
	}
}
