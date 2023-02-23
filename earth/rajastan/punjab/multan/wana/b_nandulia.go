package wana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 难菟梨NanduliaBarony struct {
	feud.BaseBarony
}

var BNandulia难菟梨 feud.Barony = &难菟梨NanduliaBarony{}

func init() {
	f := BNandulia难菟梨.(*难菟梨NanduliaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nandulia",
		TitleName: "难菟梨",
		TitleCode: "b_nandulia",
	}
}
