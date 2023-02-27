package hainaut

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢夫尔ChievresBarony struct {
	feud.BaseBarony
}

var BChievres谢夫尔 feud.Barony = &谢夫尔ChievresBarony{}

func init() {
    f := BChievres谢夫尔.(*谢夫尔ChievresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chievres",
		TitleName: "谢夫尔",
		TitleCode: "b_chievres",
	}
}
