package tunis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 突尼斯TunisBarony struct {
	feud.BaseBarony
}

var BTunis突尼斯 feud.Barony = &突尼斯TunisBarony{}

func init() {
	f := BTunis突尼斯.(*突尼斯TunisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tunis",
		TitleName: "突尼斯",
		TitleCode: "b_tunis",
	}
}
