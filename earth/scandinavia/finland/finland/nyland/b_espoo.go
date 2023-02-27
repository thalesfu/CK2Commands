package nyland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯波EspooBarony struct {
	feud.BaseBarony
}

var BEspoo埃斯波 feud.Barony = &埃斯波EspooBarony{}

func init() {
    f := BEspoo埃斯波.(*埃斯波EspooBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "espoo",
		TitleName: "埃斯波",
		TitleCode: "b_espoo",
	}
}
