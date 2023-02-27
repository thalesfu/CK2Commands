package swetaka_mandala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弃折梨耶KhijariaBarony struct {
	feud.BaseBarony
}

var BKhijaria弃折梨耶 feud.Barony = &弃折梨耶KhijariaBarony{}

func init() {
    f := BKhijaria弃折梨耶.(*弃折梨耶KhijariaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khijaria",
		TitleName: "弃折梨耶",
		TitleCode: "b_khijaria",
	}
}
