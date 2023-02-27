package kudymkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库德姆卡尔KudymkarBarony struct {
	feud.BaseBarony
}

var BKudymkar库德姆卡尔 feud.Barony = &库德姆卡尔KudymkarBarony{}

func init() {
    f := BKudymkar库德姆卡尔.(*库德姆卡尔KudymkarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kudymkar",
		TitleName: "库德姆卡尔",
		TitleCode: "b_kudymkar",
	}
}
