package venadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿骚陀AsaudahBarony struct {
	feud.BaseBarony
}

var BAsaudah阿骚陀 feud.Barony = &阿骚陀AsaudahBarony{}

func init() {
    f := BAsaudah阿骚陀.(*阿骚陀AsaudahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "asaudah",
		TitleName: "阿骚陀",
		TitleCode: "b_asaudah",
	}
}
