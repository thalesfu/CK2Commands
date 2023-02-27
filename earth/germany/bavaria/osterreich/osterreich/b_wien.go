package osterreich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维也纳WienBarony struct {
	feud.BaseBarony
}

var BWien维也纳 feud.Barony = &维也纳WienBarony{}

func init() {
    f := BWien维也纳.(*维也纳WienBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wien",
		TitleName: "维也纳",
		TitleCode: "b_wien",
	}
}
