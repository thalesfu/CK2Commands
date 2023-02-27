package plock

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普沃茨克PlockBarony struct {
	feud.BaseBarony
}

var BPlock普沃茨克 feud.Barony = &普沃茨克PlockBarony{}

func init() {
    f := BPlock普沃茨克.(*普沃茨克PlockBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plock",
		TitleName: "普沃茨克",
		TitleCode: "b_plock",
	}
}
