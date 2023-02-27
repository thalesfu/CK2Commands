package pinega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索特卡SotkaBarony struct {
	feud.BaseBarony
}

var BSotka索特卡 feud.Barony = &索特卡SotkaBarony{}

func init() {
    f := BSotka索特卡.(*索特卡SotkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sotka",
		TitleName: "索特卡",
		TitleCode: "b_sotka",
	}
}
