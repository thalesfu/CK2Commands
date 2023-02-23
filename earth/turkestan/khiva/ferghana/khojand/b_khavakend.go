package khojand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 浩罕KhavakendBarony struct {
	feud.BaseBarony
}

var BKhavakend浩罕 feud.Barony = &浩罕KhavakendBarony{}

func init() {
	f := BKhavakend浩罕.(*浩罕KhavakendBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khavakend",
		TitleName: "浩罕",
		TitleCode: "b_khavakend",
	}
}
