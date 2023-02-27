package gilgit

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉德拉尔ChitralBarony struct {
	feud.BaseBarony
}

var BChitral吉德拉尔 feud.Barony = &吉德拉尔ChitralBarony{}

func init() {
    f := BChitral吉德拉尔.(*吉德拉尔ChitralBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chitral",
		TitleName: "吉德拉尔",
		TitleCode: "b_chitral",
	}
}
