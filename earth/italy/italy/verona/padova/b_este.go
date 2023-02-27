package padova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯泰EsteBarony struct {
	feud.BaseBarony
}

var BEste埃斯泰 feud.Barony = &埃斯泰EsteBarony{}

func init() {
    f := BEste埃斯泰.(*埃斯泰EsteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "este",
		TitleName: "埃斯泰",
		TitleCode: "b_este",
	}
}
