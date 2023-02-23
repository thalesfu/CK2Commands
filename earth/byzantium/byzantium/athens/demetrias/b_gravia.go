package demetrias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉维亚GraviaBarony struct {
	feud.BaseBarony
}

var BGravia格拉维亚 feud.Barony = &格拉维亚GraviaBarony{}

func init() {
	f := BGravia格拉维亚.(*格拉维亚GraviaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gravia",
		TitleName: "格拉维亚",
		TitleCode: "b_gravia",
	}
}
