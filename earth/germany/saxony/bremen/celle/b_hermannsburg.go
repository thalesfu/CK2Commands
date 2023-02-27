package celle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黑尔曼斯堡HermannsburgBarony struct {
	feud.BaseBarony
}

var BHermannsburg黑尔曼斯堡 feud.Barony = &黑尔曼斯堡HermannsburgBarony{}

func init() {
    f := BHermannsburg黑尔曼斯堡.(*黑尔曼斯堡HermannsburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hermannsburg",
		TitleName: "黑尔曼斯堡",
		TitleCode: "b_hermannsburg",
	}
}
