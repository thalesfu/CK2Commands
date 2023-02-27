package berg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特罗特曼尼ThrotmanniBarony struct {
	feud.BaseBarony
}

var BThrotmanni特罗特曼尼 feud.Barony = &特罗特曼尼ThrotmanniBarony{}

func init() {
    f := BThrotmanni特罗特曼尼.(*特罗特曼尼ThrotmanniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "throtmanni",
		TitleName: "特罗特曼尼",
		TitleCode: "b_throtmanni",
	}
}
