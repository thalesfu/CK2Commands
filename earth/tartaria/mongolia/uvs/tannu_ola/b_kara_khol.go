package tannu_ola

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉霍尔Kara_kholBarony struct {
	feud.BaseBarony
}

var BKara_khol卡拉霍尔 feud.Barony = &卡拉霍尔Kara_kholBarony{}

func init() {
    f := BKara_khol卡拉霍尔.(*卡拉霍尔Kara_kholBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kara_khol",
		TitleName: "卡拉霍尔",
		TitleCode: "b_kara_khol",
	}
}
