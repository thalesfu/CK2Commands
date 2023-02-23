package moray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福里斯ForresBarony struct {
	feud.BaseBarony
}

var BForres福里斯 feud.Barony = &福里斯ForresBarony{}

func init() {
	f := BForres福里斯.(*福里斯ForresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "forres",
		TitleName: "福里斯",
		TitleCode: "b_forres",
	}
}
