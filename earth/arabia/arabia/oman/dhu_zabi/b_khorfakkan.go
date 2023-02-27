package dhu_zabi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 豪尔费坎KhorfakkanBarony struct {
	feud.BaseBarony
}

var BKhorfakkan豪尔费坎 feud.Barony = &豪尔费坎KhorfakkanBarony{}

func init() {
    f := BKhorfakkan豪尔费坎.(*豪尔费坎KhorfakkanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khorfakkan",
		TitleName: "豪尔费坎",
		TitleCode: "b_khorfakkan",
	}
}
