package ankober

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨卡拉SaqqaBarony struct {
	feud.BaseBarony
}

var BSaqqa萨卡拉 feud.Barony = &萨卡拉SaqqaBarony{}

func init() {
    f := BSaqqa萨卡拉.(*萨卡拉SaqqaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saqqa",
		TitleName: "萨卡拉",
		TitleCode: "b_saqqa",
	}
}
