package vas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙尔堡SarvarBarony struct {
	feud.BaseBarony
}

var BSarvar沙尔堡 feud.Barony = &沙尔堡SarvarBarony{}

func init() {
	f := BSarvar沙尔堡.(*沙尔堡SarvarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarvar",
		TitleName: "沙尔堡",
		TitleCode: "b_sarvar",
	}
}
