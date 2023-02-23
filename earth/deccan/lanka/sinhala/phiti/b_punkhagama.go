package phiti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 封佉伽摩PunkhagamaBarony struct {
	feud.BaseBarony
}

var BPunkhagama封佉伽摩 feud.Barony = &封佉伽摩PunkhagamaBarony{}

func init() {
	f := BPunkhagama封佉伽摩.(*封佉伽摩PunkhagamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "punkhagama",
		TitleName: "封佉伽摩",
		TitleCode: "b_punkhagama",
	}
}
