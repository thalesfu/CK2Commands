package daylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡纳杰KhunajBarony struct {
	feud.BaseBarony
}

var BKhunaj胡纳杰 feud.Barony = &胡纳杰KhunajBarony{}

func init() {
	f := BKhunaj胡纳杰.(*胡纳杰KhunajBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khunaj",
		TitleName: "胡纳杰",
		TitleCode: "b_khunaj",
	}
}
