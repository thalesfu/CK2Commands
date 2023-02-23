package egrisi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍比KhobiBarony struct {
	feud.BaseBarony
}

var BKhobi霍比 feud.Barony = &霍比KhobiBarony{}

func init() {
	f := BKhobi霍比.(*霍比KhobiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khobi",
		TitleName: "霍比",
		TitleCode: "b_khobi",
	}
}
