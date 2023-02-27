package aukshayts

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔塞奈AlsenaiBarony struct {
	feud.BaseBarony
}

var BAlsenai阿尔塞奈 feud.Barony = &阿尔塞奈AlsenaiBarony{}

func init() {
    f := BAlsenai阿尔塞奈.(*阿尔塞奈AlsenaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alsenai",
		TitleName: "阿尔塞奈",
		TitleCode: "b_alsenai",
	}
}
