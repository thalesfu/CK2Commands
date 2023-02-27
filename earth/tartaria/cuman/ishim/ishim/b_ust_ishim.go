package ishim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季_伊希姆Ust_ishimBarony struct {
	feud.BaseBarony
}

var BUst_ishim乌斯季_伊希姆 feud.Barony = &乌斯季_伊希姆Ust_ishimBarony{}

func init() {
    f := BUst_ishim乌斯季_伊希姆.(*乌斯季_伊希姆Ust_ishimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ust_ishim",
		TitleName: "乌斯季_伊希姆",
		TitleCode: "b_ust_ishim",
	}
}
