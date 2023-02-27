package aargau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈布斯堡HabsburgBarony struct {
	feud.BaseBarony
}

var BHabsburg哈布斯堡 feud.Barony = &哈布斯堡HabsburgBarony{}

func init() {
    f := BHabsburg哈布斯堡.(*哈布斯堡HabsburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "habsburg",
		TitleName: "哈布斯堡",
		TitleCode: "b_habsburg",
	}
}
