package novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴捷茨基BatetskyBarony struct {
	feud.BaseBarony
}

var BBatetsky巴捷茨基 feud.Barony = &巴捷茨基BatetskyBarony{}

func init() {
	f := BBatetsky巴捷茨基.(*巴捷茨基BatetskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "batetsky",
		TitleName: "巴捷茨基",
		TitleCode: "b_batetsky",
	}
}
