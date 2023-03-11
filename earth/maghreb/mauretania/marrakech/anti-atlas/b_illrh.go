package anti-atlas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安拉IllrhBarony struct {
	feud.BaseBarony
}

var BIllrh安拉 feud.Barony = &安拉IllrhBarony{}

func init() {
    f := BIllrh安拉.(*安拉IllrhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "illrh",
		TitleName: "安拉",
		TitleCode: "b_illrh",
	}
}
