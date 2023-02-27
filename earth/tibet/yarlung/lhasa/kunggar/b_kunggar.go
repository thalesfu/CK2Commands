package kunggar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 工卡KunggarBarony struct {
	feud.BaseBarony
}

var BKunggar工卡 feud.Barony = &工卡KunggarBarony{}

func init() {
    f := BKunggar工卡.(*工卡KunggarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kunggar",
		TitleName: "工卡",
		TitleCode: "b_kunggar",
	}
}
