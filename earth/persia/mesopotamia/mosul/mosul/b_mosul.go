package mosul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩苏尔MosulBarony struct {
	feud.BaseBarony
}

var BMosul摩苏尔 feud.Barony = &摩苏尔MosulBarony{}

func init() {
    f := BMosul摩苏尔.(*摩苏尔MosulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mosul",
		TitleName: "摩苏尔",
		TitleCode: "b_mosul",
	}
}
