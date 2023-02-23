package vidin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库拉KulaBarony struct {
	feud.BaseBarony
}

var BKula库拉 feud.Barony = &库拉KulaBarony{}

func init() {
	f := BKula库拉.(*库拉KulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kula",
		TitleName: "库拉",
		TitleCode: "b_kula",
	}
}
