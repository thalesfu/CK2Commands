package samarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 浩罕KhokandBarony struct {
	feud.BaseBarony
}

var BKhokand浩罕 feud.Barony = &浩罕KhokandBarony{}

func init() {
	f := BKhokand浩罕.(*浩罕KhokandBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khokand",
		TitleName: "浩罕",
		TitleCode: "b_khokand",
	}
}
