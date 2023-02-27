package najran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈萨卡特HasakatBarony struct {
	feud.BaseBarony
}

var BHasakat哈萨卡特 feud.Barony = &哈萨卡特HasakatBarony{}

func init() {
    f := BHasakat哈萨卡特.(*哈萨卡特HasakatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hasakat",
		TitleName: "哈萨卡特",
		TitleCode: "b_hasakat",
	}
}
