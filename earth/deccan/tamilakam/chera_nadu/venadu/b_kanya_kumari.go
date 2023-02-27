package venadu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 羯若鸠摩利Kanya_kumariBarony struct {
	feud.BaseBarony
}

var BKanya_kumari羯若鸠摩利 feud.Barony = &羯若鸠摩利Kanya_kumariBarony{}

func init() {
    f := BKanya_kumari羯若鸠摩利.(*羯若鸠摩利Kanya_kumariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kanya_kumari",
		TitleName: "羯若鸠摩利",
		TitleCode: "b_kanya_kumari",
	}
}
