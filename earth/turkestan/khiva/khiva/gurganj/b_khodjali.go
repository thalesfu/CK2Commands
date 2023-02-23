package gurganj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍卓里KhodjaliBarony struct {
	feud.BaseBarony
}

var BKhodjali霍卓里 feud.Barony = &霍卓里KhodjaliBarony{}

func init() {
	f := BKhodjali霍卓里.(*霍卓里KhodjaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khodjali",
		TitleName: "霍卓里",
		TitleCode: "b_khodjali",
	}
}
