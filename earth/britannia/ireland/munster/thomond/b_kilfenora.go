package thomond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔费诺拉KilfenoraBarony struct {
	feud.BaseBarony
}

var BKilfenora基尔费诺拉 feud.Barony = &基尔费诺拉KilfenoraBarony{}

func init() {
	f := BKilfenora基尔费诺拉.(*基尔费诺拉KilfenoraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kilfenora",
		TitleName: "基尔费诺拉",
		TitleCode: "b_kilfenora",
	}
}
