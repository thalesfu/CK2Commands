package gevaudan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿普谢ApchierBarony struct {
	feud.BaseBarony
}

var BApchier阿普谢 feud.Barony = &阿普谢ApchierBarony{}

func init() {
	f := BApchier阿普谢.(*阿普谢ApchierBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "apchier",
		TitleName: "阿普谢",
		TitleCode: "b_apchier",
	}
}
