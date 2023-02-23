package irgiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙赫塔ShakthaBarony struct {
	feud.BaseBarony
}

var BShaktha沙赫塔 feud.Barony = &沙赫塔ShakthaBarony{}

func init() {
	f := BShaktha沙赫塔.(*沙赫塔ShakthaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shaktha",
		TitleName: "沙赫塔",
		TitleCode: "b_shaktha",
	}
}
