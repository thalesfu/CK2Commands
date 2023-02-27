package arigh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉卜塔RabtaBarony struct {
	feud.BaseBarony
}

var BRabta拉卜塔 feud.Barony = &拉卜塔RabtaBarony{}

func init() {
    f := BRabta拉卜塔.(*拉卜塔RabtaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rabta",
		TitleName: "拉卜塔",
		TitleCode: "b_rabta",
	}
}
