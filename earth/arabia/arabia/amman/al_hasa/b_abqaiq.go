package al_hasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布盖格AbqaiqBarony struct {
	feud.BaseBarony
}

var BAbqaiq布盖格 feud.Barony = &布盖格AbqaiqBarony{}

func init() {
    f := BAbqaiq布盖格.(*布盖格AbqaiqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abqaiq",
		TitleName: "布盖格",
		TitleCode: "b_abqaiq",
	}
}
