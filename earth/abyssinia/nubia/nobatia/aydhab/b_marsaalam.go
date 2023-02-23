package aydhab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马萨阿拉姆MarsaalamBarony struct {
	feud.BaseBarony
}

var BMarsaalam马萨阿拉姆 feud.Barony = &马萨阿拉姆MarsaalamBarony{}

func init() {
	f := BMarsaalam马萨阿拉姆.(*马萨阿拉姆MarsaalamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marsaalam",
		TitleName: "马萨阿拉姆",
		TitleCode: "b_marsaalam",
	}
}
