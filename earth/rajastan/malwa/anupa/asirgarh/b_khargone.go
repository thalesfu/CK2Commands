package asirgarh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦戈讷KhargoneBarony struct {
	feud.BaseBarony
}

var BKhargone迦戈讷 feud.Barony = &迦戈讷KhargoneBarony{}

func init() {
	f := BKhargone迦戈讷.(*迦戈讷KhargoneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khargone",
		TitleName: "迦戈讷",
		TitleCode: "b_khargone",
	}
}
