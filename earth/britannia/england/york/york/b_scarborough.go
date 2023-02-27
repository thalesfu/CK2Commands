package york

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯卡布罗ScarboroughBarony struct {
	feud.BaseBarony
}

var BScarborough斯卡布罗 feud.Barony = &斯卡布罗ScarboroughBarony{}

func init() {
    f := BScarborough斯卡布罗.(*斯卡布罗ScarboroughBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "scarborough",
		TitleName: "斯卡布罗",
		TitleCode: "b_scarborough",
	}
}
