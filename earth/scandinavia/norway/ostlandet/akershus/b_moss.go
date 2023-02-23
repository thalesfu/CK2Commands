package akershus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫斯MossBarony struct {
	feud.BaseBarony
}

var BMoss莫斯 feud.Barony = &莫斯MossBarony{}

func init() {
	f := BMoss莫斯.(*莫斯MossBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moss",
		TitleName: "莫斯",
		TitleCode: "b_moss",
	}
}
