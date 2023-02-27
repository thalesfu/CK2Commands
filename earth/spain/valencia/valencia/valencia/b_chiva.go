package valencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇瓦ChivaBarony struct {
	feud.BaseBarony
}

var BChiva奇瓦 feud.Barony = &奇瓦ChivaBarony{}

func init() {
    f := BChiva奇瓦.(*奇瓦ChivaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chiva",
		TitleName: "奇瓦",
		TitleCode: "b_chiva",
	}
}
