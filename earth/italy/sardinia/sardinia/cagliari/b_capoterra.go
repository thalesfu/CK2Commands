package cagliari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡波泰拉CapoterraBarony struct {
	feud.BaseBarony
}

var BCapoterra卡波泰拉 feud.Barony = &卡波泰拉CapoterraBarony{}

func init() {
    f := BCapoterra卡波泰拉.(*卡波泰拉CapoterraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "capoterra",
		TitleName: "卡波泰拉",
		TitleCode: "b_capoterra",
	}
}
