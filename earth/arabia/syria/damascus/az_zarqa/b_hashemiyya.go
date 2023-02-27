package az_zarqa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈希姆HashemiyyaBarony struct {
	feud.BaseBarony
}

var BHashemiyya哈希姆 feud.Barony = &哈希姆HashemiyyaBarony{}

func init() {
    f := BHashemiyya哈希姆.(*哈希姆HashemiyyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hashemiyya",
		TitleName: "哈希姆",
		TitleCode: "b_hashemiyya",
	}
}
