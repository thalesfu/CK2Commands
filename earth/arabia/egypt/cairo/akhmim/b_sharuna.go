package akhmim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙鲁纳SharunaBarony struct {
	feud.BaseBarony
}

var BSharuna沙鲁纳 feud.Barony = &沙鲁纳SharunaBarony{}

func init() {
    f := BSharuna沙鲁纳.(*沙鲁纳SharunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sharuna",
		TitleName: "沙鲁纳",
		TitleCode: "b_sharuna",
	}
}
