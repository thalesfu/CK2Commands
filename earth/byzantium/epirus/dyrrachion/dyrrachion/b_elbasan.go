package dyrrachion

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃尔巴桑ElbasanBarony struct {
	feud.BaseBarony
}

var BElbasan埃尔巴桑 feud.Barony = &埃尔巴桑ElbasanBarony{}

func init() {
	f := BElbasan埃尔巴桑.(*埃尔巴桑ElbasanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elbasan",
		TitleName: "埃尔巴桑",
		TitleCode: "b_elbasan",
	}
}
