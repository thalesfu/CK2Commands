package celle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 汉诺威HannoverBarony struct {
	feud.BaseBarony
}

var BHannover汉诺威 feud.Barony = &汉诺威HannoverBarony{}

func init() {
    f := BHannover汉诺威.(*汉诺威HannoverBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hannover",
		TitleName: "汉诺威",
		TitleCode: "b_hannover",
	}
}
