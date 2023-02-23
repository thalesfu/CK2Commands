package dalarna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吞拿TunaBarony struct {
	feud.BaseBarony
}

var BTuna吞拿 feud.Barony = &吞拿TunaBarony{}

func init() {
	f := BTuna吞拿.(*吞拿TunaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tuna",
		TitleName: "吞拿",
		TitleCode: "b_tuna",
	}
}
