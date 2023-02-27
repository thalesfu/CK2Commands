package ashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿西尔_亚西尔Asir_yasirBarony struct {
	feud.BaseBarony
}

var BAsir_yasir阿西尔_亚西尔 feud.Barony = &阿西尔_亚西尔Asir_yasirBarony{}

func init() {
    f := BAsir_yasir阿西尔_亚西尔.(*阿西尔_亚西尔Asir_yasirBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asir_yasir",
		TitleName: "阿西尔_亚西尔",
		TitleCode: "b_asir_yasir",
	}
}
