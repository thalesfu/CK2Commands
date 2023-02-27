package khopyor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍皮奥尔斯克KhopyorskBarony struct {
	feud.BaseBarony
}

var BKhopyorsk霍皮奥尔斯克 feud.Barony = &霍皮奥尔斯克KhopyorskBarony{}

func init() {
    f := BKhopyorsk霍皮奥尔斯克.(*霍皮奥尔斯克KhopyorskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khopyorsk",
		TitleName: "霍皮奥尔斯克",
		TitleCode: "b_khopyorsk",
	}
}
