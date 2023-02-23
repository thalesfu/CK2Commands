package taradavadi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多罗陀婆稚TaradavadiBarony struct {
	feud.BaseBarony
}

var BTaradavadi多罗陀婆稚 feud.Barony = &多罗陀婆稚TaradavadiBarony{}

func init() {
	f := BTaradavadi多罗陀婆稚.(*多罗陀婆稚TaradavadiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taradavadi",
		TitleName: "多罗陀婆稚",
		TitleCode: "b_taradavadi",
	}
}
