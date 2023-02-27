package dhu_zabi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佐尔法JulfarBarony struct {
	feud.BaseBarony
}

var BJulfar佐尔法 feud.Barony = &佐尔法JulfarBarony{}

func init() {
    f := BJulfar佐尔法.(*佐尔法JulfarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "julfar",
		TitleName: "佐尔法",
		TitleCode: "b_julfar",
	}
}
