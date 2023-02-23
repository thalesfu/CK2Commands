package constantine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫鲁卜ElkhroubBarony struct {
	feud.BaseBarony
}

var BElkhroub赫鲁卜 feud.Barony = &赫鲁卜ElkhroubBarony{}

func init() {
	f := BElkhroub赫鲁卜.(*赫鲁卜ElkhroubBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "elkhroub",
		TitleName: "赫鲁卜",
		TitleCode: "b_elkhroub",
	}
}
