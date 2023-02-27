package munda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 周底耶ChutiaBarony struct {
	feud.BaseBarony
}

var BChutia周底耶 feud.Barony = &周底耶ChutiaBarony{}

func init() {
    f := BChutia周底耶.(*周底耶ChutiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chutia",
		TitleName: "周底耶",
		TitleCode: "b_chutia",
	}
}
