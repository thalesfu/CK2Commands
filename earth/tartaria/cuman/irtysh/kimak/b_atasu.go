package kimak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿塔苏AtasuBarony struct {
	feud.BaseBarony
}

var BAtasu阿塔苏 feud.Barony = &阿塔苏AtasuBarony{}

func init() {
    f := BAtasu阿塔苏.(*阿塔苏AtasuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atasu",
		TitleName: "阿塔苏",
		TitleCode: "b_atasu",
	}
}
