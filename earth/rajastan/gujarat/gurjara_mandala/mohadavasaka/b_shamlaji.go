package mohadavasaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨罗吉ShamlajiBarony struct {
	feud.BaseBarony
}

var BShamlaji萨罗吉 feud.Barony = &萨罗吉ShamlajiBarony{}

func init() {
    f := BShamlaji萨罗吉.(*萨罗吉ShamlajiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shamlaji",
		TitleName: "萨罗吉",
		TitleCode: "b_shamlaji",
	}
}
