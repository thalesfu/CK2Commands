package imeretia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍尼KhoniBarony struct {
	feud.BaseBarony
}

var BKhoni霍尼 feud.Barony = &霍尼KhoniBarony{}

func init() {
	f := BKhoni霍尼.(*霍尼KhoniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khoni",
		TitleName: "霍尼",
		TitleCode: "b_khoni",
	}
}
