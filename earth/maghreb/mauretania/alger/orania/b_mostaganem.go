package orania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫斯塔加纳姆MostaganemBarony struct {
	feud.BaseBarony
}

var BMostaganem莫斯塔加纳姆 feud.Barony = &莫斯塔加纳姆MostaganemBarony{}

func init() {
    f := BMostaganem莫斯塔加纳姆.(*莫斯塔加纳姆MostaganemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mostaganem",
		TitleName: "莫斯塔加纳姆",
		TitleCode: "b_mostaganem",
	}
}
