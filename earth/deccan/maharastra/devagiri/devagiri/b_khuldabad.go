package devagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔达巴德KhuldabadBarony struct {
	feud.BaseBarony
}

var BKhuldabad库尔达巴德 feud.Barony = &库尔达巴德KhuldabadBarony{}

func init() {
    f := BKhuldabad库尔达巴德.(*库尔达巴德KhuldabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khuldabad",
		TitleName: "库尔达巴德",
		TitleCode: "b_khuldabad",
	}
}
