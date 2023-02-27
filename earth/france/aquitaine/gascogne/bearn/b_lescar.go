package bearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莱斯卡尔LescarBarony struct {
	feud.BaseBarony
}

var BLescar莱斯卡尔 feud.Barony = &莱斯卡尔LescarBarony{}

func init() {
    f := BLescar莱斯卡尔.(*莱斯卡尔LescarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lescar",
		TitleName: "莱斯卡尔",
		TitleCode: "b_lescar",
	}
}
