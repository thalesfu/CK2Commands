package sussex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伦德尔ArundelBarony struct {
	feud.BaseBarony
}

var BArundel阿伦德尔 feud.Barony = &阿伦德尔ArundelBarony{}

func init() {
    f := BArundel阿伦德尔.(*阿伦德尔ArundelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "arundel",
		TitleName: "阿伦德尔",
		TitleCode: "b_arundel",
	}
}
