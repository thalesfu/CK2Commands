package pronsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格里亚兹诺耶GryaznoyeBarony struct {
	feud.BaseBarony
}

var BGryaznoye格里亚兹诺耶 feud.Barony = &格里亚兹诺耶GryaznoyeBarony{}

func init() {
	f := BGryaznoye格里亚兹诺耶.(*格里亚兹诺耶GryaznoyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gryaznoye",
		TitleName: "格里亚兹诺耶",
		TitleCode: "b_gryaznoye",
	}
}
