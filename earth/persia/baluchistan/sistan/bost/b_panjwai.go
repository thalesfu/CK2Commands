package bost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 潘杰瓦伊PanjwaiBarony struct {
	feud.BaseBarony
}

var BPanjwai潘杰瓦伊 feud.Barony = &潘杰瓦伊PanjwaiBarony{}

func init() {
	f := BPanjwai潘杰瓦伊.(*潘杰瓦伊PanjwaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "panjwai",
		TitleName: "潘杰瓦伊",
		TitleCode: "b_panjwai",
	}
}
