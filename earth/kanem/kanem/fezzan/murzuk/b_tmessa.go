package murzuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特梅萨TmessaBarony struct {
	feud.BaseBarony
}

var BTmessa特梅萨 feud.Barony = &特梅萨TmessaBarony{}

func init() {
	f := BTmessa特梅萨.(*特梅萨TmessaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tmessa",
		TitleName: "特梅萨",
		TitleCode: "b_tmessa",
	}
}
