package anglesey

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伯弗劳AberffrawBarony struct {
	feud.BaseBarony
}

var BAberffraw阿伯弗劳 feud.Barony = &阿伯弗劳AberffrawBarony{}

func init() {
    f := BAberffraw阿伯弗劳.(*阿伯弗劳AberffrawBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aberffraw",
		TitleName: "阿伯弗劳",
		TitleCode: "b_aberffraw",
	}
}
