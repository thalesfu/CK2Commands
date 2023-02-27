package alania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔马拉舍尼TamarasheniBarony struct {
	feud.BaseBarony
}

var BTamarasheni塔马拉舍尼 feud.Barony = &塔马拉舍尼TamarasheniBarony{}

func init() {
    f := BTamarasheni塔马拉舍尼.(*塔马拉舍尼TamarasheniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tamarasheni",
		TitleName: "塔马拉舍尼",
		TitleCode: "b_tamarasheni",
	}
}
