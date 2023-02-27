package karachev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳雷什基诺NaryshkinoBarony struct {
	feud.BaseBarony
}

var BNaryshkino纳雷什基诺 feud.Barony = &纳雷什基诺NaryshkinoBarony{}

func init() {
    f := BNaryshkino纳雷什基诺.(*纳雷什基诺NaryshkinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naryshkino",
		TitleName: "纳雷什基诺",
		TitleCode: "b_naryshkino",
	}
}
