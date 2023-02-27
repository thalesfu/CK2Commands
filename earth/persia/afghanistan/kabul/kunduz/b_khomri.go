package kunduz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡姆里KhomriBarony struct {
	feud.BaseBarony
}

var BKhomri胡姆里 feud.Barony = &胡姆里KhomriBarony{}

func init() {
    f := BKhomri胡姆里.(*胡姆里KhomriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khomri",
		TitleName: "胡姆里",
		TitleCode: "b_khomri",
	}
}
