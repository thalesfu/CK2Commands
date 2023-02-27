package dunhuang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 悬泉置XuanquanzhiBarony struct {
	feud.BaseBarony
}

var BXuanquanzhi悬泉置 feud.Barony = &悬泉置XuanquanzhiBarony{}

func init() {
    f := BXuanquanzhi悬泉置.(*悬泉置XuanquanzhiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xuanquanzhi",
		TitleName: "悬泉置",
		TitleCode: "b_xuanquanzhi",
	}
}
