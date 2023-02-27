package ashir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿西尔_班亚Asir_banyaBarony struct {
	feud.BaseBarony
}

var BAsir_banya阿西尔_班亚 feud.Barony = &阿西尔_班亚Asir_banyaBarony{}

func init() {
    f := BAsir_banya阿西尔_班亚.(*阿西尔_班亚Asir_banyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asir_banya",
		TitleName: "阿西尔_班亚",
		TitleCode: "b_asir_banya",
	}
}
