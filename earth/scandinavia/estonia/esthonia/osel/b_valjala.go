package osel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦利亚拉ValjalaBarony struct {
	feud.BaseBarony
}

var BValjala瓦利亚拉 feud.Barony = &瓦利亚拉ValjalaBarony{}

func init() {
	f := BValjala瓦利亚拉.(*瓦利亚拉ValjalaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "valjala",
		TitleName: "瓦利亚拉",
		TitleCode: "b_valjala",
	}
}
