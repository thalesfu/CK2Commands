package alcacer_do_sal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣地亚哥杜卡森SantiagodocacemBarony struct {
	feud.BaseBarony
}

var BSantiagodocacem圣地亚哥杜卡森 feud.Barony = &圣地亚哥杜卡森SantiagodocacemBarony{}

func init() {
    f := BSantiagodocacem圣地亚哥杜卡森.(*圣地亚哥杜卡森SantiagodocacemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santiagodocacem",
		TitleName: "圣地亚哥杜卡森",
		TitleCode: "b_santiagodocacem",
	}
}
