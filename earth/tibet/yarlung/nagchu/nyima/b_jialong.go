package nyima

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加龙JialongBarony struct {
	feud.BaseBarony
}

var BJialong加龙 feud.Barony = &加龙JialongBarony{}

func init() {
	f := BJialong加龙.(*加龙JialongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jialong",
		TitleName: "加龙",
		TitleCode: "b_jialong",
	}
}
