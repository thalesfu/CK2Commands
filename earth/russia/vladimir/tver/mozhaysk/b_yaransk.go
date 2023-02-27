package mozhaysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚兰斯克YaranskBarony struct {
	feud.BaseBarony
}

var BYaransk亚兰斯克 feud.Barony = &亚兰斯克YaranskBarony{}

func init() {
    f := BYaransk亚兰斯克.(*亚兰斯克YaranskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yaransk",
		TitleName: "亚兰斯克",
		TitleCode: "b_yaransk",
	}
}
