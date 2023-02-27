package vashka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦什卡VashkaBarony struct {
	feud.BaseBarony
}

var BVashka瓦什卡 feud.Barony = &瓦什卡VashkaBarony{}

func init() {
    f := BVashka瓦什卡.(*瓦什卡VashkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vashka",
		TitleName: "瓦什卡",
		TitleCode: "b_vashka",
	}
}
