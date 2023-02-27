package nangqen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尕涌GayongBarony struct {
	feud.BaseBarony
}

var BGayong尕涌 feud.Barony = &尕涌GayongBarony{}

func init() {
    f := BGayong尕涌.(*尕涌GayongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gayong",
		TitleName: "尕涌",
		TitleCode: "b_gayong",
	}
}
