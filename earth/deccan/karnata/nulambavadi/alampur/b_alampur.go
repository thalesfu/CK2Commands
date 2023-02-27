package alampur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿蓝补罗AlampurBarony struct {
	feud.BaseBarony
}

var BAlampur阿蓝补罗 feud.Barony = &阿蓝补罗AlampurBarony{}

func init() {
    f := BAlampur阿蓝补罗.(*阿蓝补罗AlampurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alampur",
		TitleName: "阿蓝补罗",
		TitleCode: "b_alampur",
	}
}
