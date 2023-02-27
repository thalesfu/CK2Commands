package akroinon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿帕美亚Apamea_akroinonBarony struct {
	feud.BaseBarony
}

var BApamea_akroinon阿帕美亚 feud.Barony = &阿帕美亚Apamea_akroinonBarony{}

func init() {
    f := BApamea_akroinon阿帕美亚.(*阿帕美亚Apamea_akroinonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "apamea_akroinon",
		TitleName: "阿帕美亚",
		TitleCode: "b_apamea_akroinon",
	}
}
