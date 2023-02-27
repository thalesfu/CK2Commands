package delingha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蓄集XujiBarony struct {
	feud.BaseBarony
}

var BXuji蓄集 feud.Barony = &蓄集XujiBarony{}

func init() {
    f := BXuji蓄集.(*蓄集XujiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "xuji",
		TitleName: "蓄集",
		TitleCode: "b_xuji",
	}
}
