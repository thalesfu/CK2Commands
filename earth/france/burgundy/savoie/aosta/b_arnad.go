package aosta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔纳德ArnadBarony struct {
	feud.BaseBarony
}

var BArnad阿尔纳德 feud.Barony = &阿尔纳德ArnadBarony{}

func init() {
	f := BArnad阿尔纳德.(*阿尔纳德ArnadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arnad",
		TitleName: "阿尔纳德",
		TitleCode: "b_arnad",
	}
}
