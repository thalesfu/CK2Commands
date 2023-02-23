package palmyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希尔拜特KhirbatBarony struct {
	feud.BaseBarony
}

var BKhirbat希尔拜特 feud.Barony = &希尔拜特KhirbatBarony{}

func init() {
	f := BKhirbat希尔拜特.(*希尔拜特KhirbatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khirbat",
		TitleName: "希尔拜特",
		TitleCode: "b_khirbat",
	}
}
