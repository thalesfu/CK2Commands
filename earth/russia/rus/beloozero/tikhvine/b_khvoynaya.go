package tikhvine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫沃伊纳亚KhvoynayaBarony struct {
	feud.BaseBarony
}

var BKhvoynaya赫沃伊纳亚 feud.Barony = &赫沃伊纳亚KhvoynayaBarony{}

func init() {
    f := BKhvoynaya赫沃伊纳亚.(*赫沃伊纳亚KhvoynayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khvoynaya",
		TitleName: "赫沃伊纳亚",
		TitleCode: "b_khvoynaya",
	}
}
