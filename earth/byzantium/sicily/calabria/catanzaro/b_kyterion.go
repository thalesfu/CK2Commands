package catanzaro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基泰里翁KyterionBarony struct {
	feud.BaseBarony
}

var BKyterion基泰里翁 feud.Barony = &基泰里翁KyterionBarony{}

func init() {
	f := BKyterion基泰里翁.(*基泰里翁KyterionBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kyterion",
		TitleName: "基泰里翁",
		TitleCode: "b_kyterion",
	}
}
