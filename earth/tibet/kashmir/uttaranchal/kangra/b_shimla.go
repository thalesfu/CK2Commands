package kangra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西姆拉ShimlaBarony struct {
	feud.BaseBarony
}

var BShimla西姆拉 feud.Barony = &西姆拉ShimlaBarony{}

func init() {
    f := BShimla西姆拉.(*西姆拉ShimlaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shimla",
		TitleName: "西姆拉",
		TitleCode: "b_shimla",
	}
}
