package kent

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多佛尔DoverBarony struct {
	feud.BaseBarony
}

var BDover多佛尔 feud.Barony = &多佛尔DoverBarony{}

func init() {
	f := BDover多佛尔.(*多佛尔DoverBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dover",
		TitleName: "多佛尔",
		TitleCode: "b_dover",
	}
}
