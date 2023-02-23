package kerak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔哈雷塞特KirhasesetBarony struct {
	feud.BaseBarony
}

var BKirhaseset基尔哈雷塞特 feud.Barony = &基尔哈雷塞特KirhasesetBarony{}

func init() {
	f := BKirhaseset基尔哈雷塞特.(*基尔哈雷塞特KirhasesetBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirhaseset",
		TitleName: "基尔哈雷塞特",
		TitleCode: "b_kirhaseset",
	}
}
