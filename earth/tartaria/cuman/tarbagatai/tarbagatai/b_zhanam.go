package tarbagatai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎纳姆ZhanamBarony struct {
	feud.BaseBarony
}

var BZhanam扎纳姆 feud.Barony = &扎纳姆ZhanamBarony{}

func init() {
	f := BZhanam扎纳姆.(*扎纳姆ZhanamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhanam",
		TitleName: "扎纳姆",
		TitleCode: "b_zhanam",
	}
}
