package quena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈拉夫KhalafBarony struct {
	feud.BaseBarony
}

var BKhalaf哈拉夫 feud.Barony = &哈拉夫KhalafBarony{}

func init() {
	f := BKhalaf哈拉夫.(*哈拉夫KhalafBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khalaf",
		TitleName: "哈拉夫",
		TitleCode: "b_khalaf",
	}
}
