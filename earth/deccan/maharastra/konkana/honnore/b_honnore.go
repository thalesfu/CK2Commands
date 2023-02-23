package honnore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡奴梨HonnoreBarony struct {
	feud.BaseBarony
}

var BHonnore胡奴梨 feud.Barony = &胡奴梨HonnoreBarony{}

func init() {
	f := BHonnore胡奴梨.(*胡奴梨HonnoreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "honnore",
		TitleName: "胡奴梨",
		TitleCode: "b_honnore",
	}
}
