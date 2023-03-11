package zvenyhorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹韦尼戈罗德ZvenyhorodBarony struct {
	feud.BaseBarony
}

var BZvenyhorod兹韦尼戈罗德 feud.Barony = &兹韦尼戈罗德ZvenyhorodBarony{}

func init() {
    f := BZvenyhorod兹韦尼戈罗德.(*兹韦尼戈罗德ZvenyhorodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zvenyhorod",
		TitleName: "兹韦尼戈罗德",
		TitleCode: "b_zvenyhorod",
	}
}
