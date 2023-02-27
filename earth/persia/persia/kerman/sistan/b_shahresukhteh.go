package sistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙赫尔苏赫特ShahresukhtehBarony struct {
	feud.BaseBarony
}

var BShahresukhteh沙赫尔苏赫特 feud.Barony = &沙赫尔苏赫特ShahresukhtehBarony{}

func init() {
    f := BShahresukhteh沙赫尔苏赫特.(*沙赫尔苏赫特ShahresukhtehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shahresukhteh",
		TitleName: "沙赫尔苏赫特",
		TitleCode: "b_shahresukhteh",
	}
}
