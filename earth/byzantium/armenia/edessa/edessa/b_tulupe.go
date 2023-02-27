package edessa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图鲁匹TulupeBarony struct {
	feud.BaseBarony
}

var BTulupe图鲁匹 feud.Barony = &图鲁匹TulupeBarony{}

func init() {
    f := BTulupe图鲁匹.(*图鲁匹TulupeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tulupe",
		TitleName: "图鲁匹",
		TitleCode: "b_tulupe",
	}
}
