package shemakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈奇马兹KhachmazBarony struct {
	feud.BaseBarony
}

var BKhachmaz哈奇马兹 feud.Barony = &哈奇马兹KhachmazBarony{}

func init() {
	f := BKhachmaz哈奇马兹.(*哈奇马兹KhachmazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khachmaz",
		TitleName: "哈奇马兹",
		TitleCode: "b_khachmaz",
	}
}
