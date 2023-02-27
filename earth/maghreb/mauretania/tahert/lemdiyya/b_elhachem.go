package lemdiyya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈希姆ElhachemBarony struct {
	feud.BaseBarony
}

var BElhachem哈希姆 feud.Barony = &哈希姆ElhachemBarony{}

func init() {
    f := BElhachem哈希姆.(*哈希姆ElhachemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elhachem",
		TitleName: "哈希姆",
		TitleCode: "b_elhachem",
	}
}
