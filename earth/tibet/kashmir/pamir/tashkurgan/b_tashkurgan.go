package tashkurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔什库尔干TashkurganBarony struct {
	feud.BaseBarony
}

var BTashkurgan塔什库尔干 feud.Barony = &塔什库尔干TashkurganBarony{}

func init() {
    f := BTashkurgan塔什库尔干.(*塔什库尔干TashkurganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tashkurgan",
		TitleName: "塔什库尔干",
		TitleCode: "b_tashkurgan",
	}
}
