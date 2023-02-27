package vychegda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁奇RuchBarony struct {
	feud.BaseBarony
}

var BRuch鲁奇 feud.Barony = &鲁奇RuchBarony{}

func init() {
    f := BRuch鲁奇.(*鲁奇RuchBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ruch",
		TitleName: "鲁奇",
		TitleCode: "b_ruch",
	}
}
