package onega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季_奥涅加Ust_onegaBarony struct {
	feud.BaseBarony
}

var BUst_onega乌斯季_奥涅加 feud.Barony = &乌斯季_奥涅加Ust_onegaBarony{}

func init() {
    f := BUst_onega乌斯季_奥涅加.(*乌斯季_奥涅加Ust_onegaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ust_onega",
		TitleName: "乌斯季_奥涅加",
		TitleCode: "b_ust-onega",
	}
}
