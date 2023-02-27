package kodalaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔切尔TalcherBarony struct {
	feud.BaseBarony
}

var BTalcher塔尔切尔 feud.Barony = &塔尔切尔TalcherBarony{}

func init() {
    f := BTalcher塔尔切尔.(*塔尔切尔TalcherBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talcher",
		TitleName: "塔尔切尔",
		TitleCode: "b_talcher",
	}
}
