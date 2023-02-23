package iarmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯凯利格SkelligBarony struct {
	feud.BaseBarony
}

var BSkellig斯凯利格 feud.Barony = &斯凯利格SkelligBarony{}

func init() {
	f := BSkellig斯凯利格.(*斯凯利格SkelligBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "skellig",
		TitleName: "斯凯利格",
		TitleCode: "b_skellig",
	}
}
