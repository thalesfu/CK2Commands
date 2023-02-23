package beshbaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别失八里BeshbalikBarony struct {
	feud.BaseBarony
}

var BBeshbalik别失八里 feud.Barony = &别失八里BeshbalikBarony{}

func init() {
	f := BBeshbalik别失八里.(*别失八里BeshbalikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "beshbalik",
		TitleName: "别失八里",
		TitleCode: "b_beshbalik",
	}
}
