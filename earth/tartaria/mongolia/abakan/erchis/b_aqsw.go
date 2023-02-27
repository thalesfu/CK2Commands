package erchis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克苏AqswBarony struct {
	feud.BaseBarony
}

var BAqsw阿克苏 feud.Barony = &阿克苏AqswBarony{}

func init() {
    f := BAqsw阿克苏.(*阿克苏AqswBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aqsw",
		TitleName: "阿克苏",
		TitleCode: "b_aqsw",
	}
}
