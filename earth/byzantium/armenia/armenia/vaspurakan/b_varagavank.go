package vaspurakan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦拉格凡克VaragavankBarony struct {
	feud.BaseBarony
}

var BVaragavank瓦拉格凡克 feud.Barony = &瓦拉格凡克VaragavankBarony{}

func init() {
	f := BVaragavank瓦拉格凡克.(*瓦拉格凡克VaragavankBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "varagavank",
		TitleName: "瓦拉格凡克",
		TitleCode: "b_varagavank",
	}
}
