package mordva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨兰斯克SaranskBarony struct {
	feud.BaseBarony
}

var BSaransk萨兰斯克 feud.Barony = &萨兰斯克SaranskBarony{}

func init() {
    f := BSaransk萨兰斯克.(*萨兰斯克SaranskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saransk",
		TitleName: "萨兰斯克",
		TitleCode: "b_saransk",
	}
}
