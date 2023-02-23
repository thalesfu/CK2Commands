package kaisereia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔拉斯TalasBarony struct {
	feud.BaseBarony
}

var BTalas塔拉斯 feud.Barony = &塔拉斯TalasBarony{}

func init() {
	f := BTalas塔拉斯.(*塔拉斯TalasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "talas",
		TitleName: "塔拉斯",
		TitleCode: "b_talas",
	}
}
