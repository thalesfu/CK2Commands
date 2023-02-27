package socotra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加希乌QashiuBarony struct {
	feud.BaseBarony
}

var BQashiu加希乌 feud.Barony = &加希乌QashiuBarony{}

func init() {
    f := BQashiu加希乌.(*加希乌QashiuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qashiu",
		TitleName: "加希乌",
		TitleCode: "b_qashiu",
	}
}
