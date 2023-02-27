package guryev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿特劳AtyrauBarony struct {
	feud.BaseBarony
}

var BAtyrau阿特劳 feud.Barony = &阿特劳AtyrauBarony{}

func init() {
    f := BAtyrau阿特劳.(*阿特劳AtyrauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atyrau",
		TitleName: "阿特劳",
		TitleCode: "b_atyrau",
	}
}
