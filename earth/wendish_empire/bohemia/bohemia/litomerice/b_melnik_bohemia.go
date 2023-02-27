package litomerice

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅尔尼克Melnik_bohemiaBarony struct {
	feud.BaseBarony
}

var BMelnik_bohemia梅尔尼克 feud.Barony = &梅尔尼克Melnik_bohemiaBarony{}

func init() {
    f := BMelnik_bohemia梅尔尼克.(*梅尔尼克Melnik_bohemiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "melnik_bohemia",
		TitleName: "梅尔尼克",
		TitleCode: "b_melnik_bohemia",
	}
}
