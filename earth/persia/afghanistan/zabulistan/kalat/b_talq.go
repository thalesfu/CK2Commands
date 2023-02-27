package kalat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉喀TalqBarony struct {
	feud.BaseBarony
}

var BTalq塔拉喀 feud.Barony = &塔拉喀TalqBarony{}

func init() {
    f := BTalq塔拉喀.(*塔拉喀TalqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talq",
		TitleName: "塔拉喀",
		TitleCode: "b_talq",
	}
}
