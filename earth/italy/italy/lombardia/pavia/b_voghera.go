package pavia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃盖拉VogheraBarony struct {
	feud.BaseBarony
}

var BVoghera沃盖拉 feud.Barony = &沃盖拉VogheraBarony{}

func init() {
	f := BVoghera沃盖拉.(*沃盖拉VogheraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "voghera",
		TitleName: "沃盖拉",
		TitleCode: "b_voghera",
	}
}
