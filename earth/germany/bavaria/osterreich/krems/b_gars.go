package krems

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加尔斯GarsBarony struct {
	feud.BaseBarony
}

var BGars加尔斯 feud.Barony = &加尔斯GarsBarony{}

func init() {
    f := BGars加尔斯.(*加尔斯GarsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gars",
		TitleName: "加尔斯",
		TitleCode: "b_gars",
	}
}
