package kusinagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帝利那跋陀TrinabadBarony struct {
	feud.BaseBarony
}

var BTrinabad帝利那跋陀 feud.Barony = &帝利那跋陀TrinabadBarony{}

func init() {
	f := BTrinabad帝利那跋陀.(*帝利那跋陀TrinabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trinabad",
		TitleName: "帝利那跋陀",
		TitleCode: "b_trinabad",
	}
}
