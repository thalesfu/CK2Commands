package sutai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 呼尔明KhulmiinBarony struct {
	feud.BaseBarony
}

var BKhulmiin呼尔明 feud.Barony = &呼尔明KhulmiinBarony{}

func init() {
    f := BKhulmiin呼尔明.(*呼尔明KhulmiinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khulmiin",
		TitleName: "呼尔明",
		TitleCode: "b_khulmiin",
	}
}
