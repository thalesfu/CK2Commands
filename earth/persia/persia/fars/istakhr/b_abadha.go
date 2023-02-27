package istakhr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿巴扎AbadhaBarony struct {
	feud.BaseBarony
}

var BAbadha阿巴扎 feud.Barony = &阿巴扎AbadhaBarony{}

func init() {
    f := BAbadha阿巴扎.(*阿巴扎AbadhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abadha",
		TitleName: "阿巴扎",
		TitleCode: "b_abadha",
	}
}
