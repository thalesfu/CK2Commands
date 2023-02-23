package achaia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶拉基GerakiBarony struct {
	feud.BaseBarony
}

var BGeraki耶拉基 feud.Barony = &耶拉基GerakiBarony{}

func init() {
	f := BGeraki耶拉基.(*耶拉基GerakiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "geraki",
		TitleName: "耶拉基",
		TitleCode: "b_geraki",
	}
}
