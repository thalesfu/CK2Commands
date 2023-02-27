package herakleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽菲洛波利ZephyropoliBarony struct {
	feud.BaseBarony
}

var BZephyropoli泽菲洛波利 feud.Barony = &泽菲洛波利ZephyropoliBarony{}

func init() {
    f := BZephyropoli泽菲洛波利.(*泽菲洛波利ZephyropoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zephyropoli",
		TitleName: "泽菲洛波利",
		TitleCode: "b_zephyropoli",
	}
}
