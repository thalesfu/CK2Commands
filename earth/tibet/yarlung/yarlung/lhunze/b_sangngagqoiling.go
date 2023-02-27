package lhunze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 三安曲林SangngagqoilingBarony struct {
	feud.BaseBarony
}

var BSangngagqoiling三安曲林 feud.Barony = &三安曲林SangngagqoilingBarony{}

func init() {
    f := BSangngagqoiling三安曲林.(*三安曲林SangngagqoilingBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sangngagqoiling",
		TitleName: "三安曲林",
		TitleCode: "b_sangngagqoiling",
	}
}
