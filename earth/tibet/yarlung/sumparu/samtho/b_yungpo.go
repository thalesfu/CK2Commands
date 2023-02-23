package samtho

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 永波YungpoBarony struct {
	feud.BaseBarony
}

var BYungpo永波 feud.Barony = &永波YungpoBarony{}

func init() {
	f := BYungpo永波.(*永波YungpoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yungpo",
		TitleName: "永波",
		TitleCode: "b_yungpo",
	}
}
