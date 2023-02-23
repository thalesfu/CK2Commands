package jiuquan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 酒泉JiuquanBarony struct {
	feud.BaseBarony
}

var BJiuquan酒泉 feud.Barony = &酒泉JiuquanBarony{}

func init() {
	f := BJiuquan酒泉.(*酒泉JiuquanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jiuquan",
		TitleName: "酒泉",
		TitleCode: "b_jiuquan",
	}
}
