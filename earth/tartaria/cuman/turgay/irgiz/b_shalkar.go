package irgiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙尔卡尔ShalkarBarony struct {
	feud.BaseBarony
}

var BShalkar沙尔卡尔 feud.Barony = &沙尔卡尔ShalkarBarony{}

func init() {
	f := BShalkar沙尔卡尔.(*沙尔卡尔ShalkarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shalkar",
		TitleName: "沙尔卡尔",
		TitleCode: "b_shalkar",
	}
}
