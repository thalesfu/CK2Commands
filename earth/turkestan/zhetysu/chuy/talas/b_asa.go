package talas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿萨AsaBarony struct {
	feud.BaseBarony
}

var BAsa阿萨 feud.Barony = &阿萨AsaBarony{}

func init() {
	f := BAsa阿萨.(*阿萨AsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asa",
		TitleName: "阿萨",
		TitleCode: "b_asa",
	}
}
