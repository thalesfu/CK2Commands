package orkhon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡夏特KhushaatBarony struct {
	feud.BaseBarony
}

var BKhushaat胡夏特 feud.Barony = &胡夏特KhushaatBarony{}

func init() {
    f := BKhushaat胡夏特.(*胡夏特KhushaatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khushaat",
		TitleName: "胡夏特",
		TitleCode: "b_khushaat",
	}
}
