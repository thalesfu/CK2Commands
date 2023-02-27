package taizz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔伊兹TaizzBarony struct {
	feud.BaseBarony
}

var BTaizz塔伊兹 feud.Barony = &塔伊兹TaizzBarony{}

func init() {
    f := BTaizz塔伊兹.(*塔伊兹TaizzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "taizz",
		TitleName: "塔伊兹",
		TitleCode: "b_taizz",
	}
}
