package trondelag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特隆赫姆TrondheimBarony struct {
	feud.BaseBarony
}

var BTrondheim特隆赫姆 feud.Barony = &特隆赫姆TrondheimBarony{}

func init() {
    f := BTrondheim特隆赫姆.(*特隆赫姆TrondheimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "trondheim",
		TitleName: "特隆赫姆",
		TitleCode: "b_trondheim",
	}
}
