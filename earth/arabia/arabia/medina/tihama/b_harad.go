package tihama

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈拉德HaradBarony struct {
	feud.BaseBarony
}

var BHarad哈拉德 feud.Barony = &哈拉德HaradBarony{}

func init() {
	f := BHarad哈拉德.(*哈拉德HaradBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "harad",
		TitleName: "哈拉德",
		TitleCode: "b_harad",
	}
}
