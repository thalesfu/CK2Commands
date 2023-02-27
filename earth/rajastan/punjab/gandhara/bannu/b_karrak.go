package bannu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉克KarrakBarony struct {
	feud.BaseBarony
}

var BKarrak卡拉克 feud.Barony = &卡拉克KarrakBarony{}

func init() {
    f := BKarrak卡拉克.(*卡拉克KarrakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karrak",
		TitleName: "卡拉克",
		TitleCode: "b_karrak",
	}
}
