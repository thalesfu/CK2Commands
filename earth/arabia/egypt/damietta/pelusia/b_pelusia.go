package pelusia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 廷尼斯PelusiaBarony struct {
	feud.BaseBarony
}

var BPelusia廷尼斯 feud.Barony = &廷尼斯PelusiaBarony{}

func init() {
	f := BPelusia廷尼斯.(*廷尼斯PelusiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pelusia",
		TitleName: "廷尼斯",
		TitleCode: "b_pelusia",
	}
}
