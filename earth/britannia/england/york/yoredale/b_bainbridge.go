package yoredale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 班布里奇BainbridgeBarony struct {
	feud.BaseBarony
}

var BBainbridge班布里奇 feud.Barony = &班布里奇BainbridgeBarony{}

func init() {
    f := BBainbridge班布里奇.(*班布里奇BainbridgeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bainbridge",
		TitleName: "班布里奇",
		TitleCode: "b_bainbridge",
	}
}
