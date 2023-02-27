package hellas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔克利尼亚MarkryniaBarony struct {
	feud.BaseBarony
}

var BMarkrynia马尔克利尼亚 feud.Barony = &马尔克利尼亚MarkryniaBarony{}

func init() {
    f := BMarkrynia马尔克利尼亚.(*马尔克利尼亚MarkryniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "markrynia",
		TitleName: "马尔克利尼亚",
		TitleCode: "b_markrynia",
	}
}
