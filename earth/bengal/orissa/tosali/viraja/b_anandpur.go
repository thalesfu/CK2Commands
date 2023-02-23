package viraja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿难陀补罗AnandpurBarony struct {
	feud.BaseBarony
}

var BAnandpur阿难陀补罗 feud.Barony = &阿难陀补罗AnandpurBarony{}

func init() {
	f := BAnandpur阿难陀补罗.(*阿难陀补罗AnandpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anandpur",
		TitleName: "阿难陀补罗",
		TitleCode: "b_anandpur",
	}
}
