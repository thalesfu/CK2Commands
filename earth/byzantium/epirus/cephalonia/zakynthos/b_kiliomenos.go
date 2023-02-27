package zakynthos

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基利奥梅诺斯KiliomenosBarony struct {
	feud.BaseBarony
}

var BKiliomenos基利奥梅诺斯 feud.Barony = &基利奥梅诺斯KiliomenosBarony{}

func init() {
    f := BKiliomenos基利奥梅诺斯.(*基利奥梅诺斯KiliomenosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiliomenos",
		TitleName: "基利奥梅诺斯",
		TitleCode: "b_kiliomenos",
	}
}
