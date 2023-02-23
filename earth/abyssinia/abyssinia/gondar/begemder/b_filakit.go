package begemder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲拉基特FilakitBarony struct {
	feud.BaseBarony
}

var BFilakit菲拉基特 feud.Barony = &菲拉基特FilakitBarony{}

func init() {
	f := BFilakit菲拉基特.(*菲拉基特FilakitBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "filakit",
		TitleName: "菲拉基特",
		TitleCode: "b_filakit",
	}
}
