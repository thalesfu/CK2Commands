package zhetysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库拉克苏KuraksuBarony struct {
	feud.BaseBarony
}

var BKuraksu库拉克苏 feud.Barony = &库拉克苏KuraksuBarony{}

func init() {
	f := BKuraksu库拉克苏.(*库拉克苏KuraksuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuraksu",
		TitleName: "库拉克苏",
		TitleCode: "b_kuraksu",
	}
}
