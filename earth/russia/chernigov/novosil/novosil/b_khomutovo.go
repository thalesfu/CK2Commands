package novosil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍穆托沃KhomutovoBarony struct {
	feud.BaseBarony
}

var BKhomutovo霍穆托沃 feud.Barony = &霍穆托沃KhomutovoBarony{}

func init() {
	f := BKhomutovo霍穆托沃.(*霍穆托沃KhomutovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khomutovo",
		TitleName: "霍穆托沃",
		TitleCode: "b_khomutovo",
	}
}
