package meissen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦廷WettinBarony struct {
	feud.BaseBarony
}

var BWettin韦廷 feud.Barony = &韦廷WettinBarony{}

func init() {
    f := BWettin韦廷.(*韦廷WettinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wettin",
		TitleName: "韦廷",
		TitleCode: "b_wettin",
	}
}
