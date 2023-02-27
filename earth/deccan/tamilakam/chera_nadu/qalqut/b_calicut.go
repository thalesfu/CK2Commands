package qalqut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡利卡特CalicutBarony struct {
	feud.BaseBarony
}

var BCalicut卡利卡特 feud.Barony = &卡利卡特CalicutBarony{}

func init() {
    f := BCalicut卡利卡特.(*卡利卡特CalicutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "calicut",
		TitleName: "卡利卡特",
		TitleCode: "b_calicut",
	}
}
