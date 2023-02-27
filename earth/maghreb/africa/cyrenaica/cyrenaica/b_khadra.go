package cyrenaica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡哈达拉KhadraBarony struct {
	feud.BaseBarony
}

var BKhadra卡哈达拉 feud.Barony = &卡哈达拉KhadraBarony{}

func init() {
    f := BKhadra卡哈达拉.(*卡哈达拉KhadraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khadra",
		TitleName: "卡哈达拉",
		TitleCode: "b_khadra",
	}
}
