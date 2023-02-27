package kusinagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿蓝补罗AlampuraBarony struct {
	feud.BaseBarony
}

var BAlampura阿蓝补罗 feud.Barony = &阿蓝补罗AlampuraBarony{}

func init() {
    f := BAlampura阿蓝补罗.(*阿蓝补罗AlampuraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alampura",
		TitleName: "阿蓝补罗",
		TitleCode: "b_alampura",
	}
}
