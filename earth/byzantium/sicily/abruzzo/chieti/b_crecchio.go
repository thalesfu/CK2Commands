package chieti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克雷基奥CrecchioBarony struct {
	feud.BaseBarony
}

var BCrecchio克雷基奥 feud.Barony = &克雷基奥CrecchioBarony{}

func init() {
	f := BCrecchio克雷基奥.(*克雷基奥CrecchioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "crecchio",
		TitleName: "克雷基奥",
		TitleCode: "b_crecchio",
	}
}
