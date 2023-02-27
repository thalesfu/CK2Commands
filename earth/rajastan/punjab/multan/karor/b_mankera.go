package karor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 门盖拉MankeraBarony struct {
	feud.BaseBarony
}

var BMankera门盖拉 feud.Barony = &门盖拉MankeraBarony{}

func init() {
    f := BMankera门盖拉.(*门盖拉MankeraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mankera",
		TitleName: "门盖拉",
		TitleCode: "b_mankera",
	}
}
