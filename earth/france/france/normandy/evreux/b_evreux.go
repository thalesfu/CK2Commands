package evreux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃夫勒EvreuxBarony struct {
	feud.BaseBarony
}

var BEvreux埃夫勒 feud.Barony = &埃夫勒EvreuxBarony{}

func init() {
	f := BEvreux埃夫勒.(*埃夫勒EvreuxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "evreux",
		TitleName: "埃夫勒",
		TitleCode: "b_evreux",
	}
}
