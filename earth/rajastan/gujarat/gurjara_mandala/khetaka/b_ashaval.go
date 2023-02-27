package khetaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿舍伐罗AshavalBarony struct {
	feud.BaseBarony
}

var BAshaval阿舍伐罗 feud.Barony = &阿舍伐罗AshavalBarony{}

func init() {
    f := BAshaval阿舍伐罗.(*阿舍伐罗AshavalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ashaval",
		TitleName: "阿舍伐罗",
		TitleCode: "b_ashaval",
	}
}
