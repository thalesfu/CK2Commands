package fejer

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞克萨德SzekszardBarony struct {
	feud.BaseBarony
}

var BSzekszard塞克萨德 feud.Barony = &塞克萨德SzekszardBarony{}

func init() {
    f := BSzekszard塞克萨德.(*塞克萨德SzekszardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "szekszard",
		TitleName: "塞克萨德",
		TitleCode: "b_szekszard",
	}
}
