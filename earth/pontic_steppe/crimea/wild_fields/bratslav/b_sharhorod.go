package bratslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙尔霍罗德SharhorodBarony struct {
	feud.BaseBarony
}

var BSharhorod沙尔霍罗德 feud.Barony = &沙尔霍罗德SharhorodBarony{}

func init() {
    f := BSharhorod沙尔霍罗德.(*沙尔霍罗德SharhorodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sharhorod",
		TitleName: "沙尔霍罗德",
		TitleCode: "b_sharhorod",
	}
}
