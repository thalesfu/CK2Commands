package braganza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙维什ChavesBarony struct {
	feud.BaseBarony
}

var BChaves沙维什 feud.Barony = &沙维什ChavesBarony{}

func init() {
	f := BChaves沙维什.(*沙维什ChavesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chaves",
		TitleName: "沙维什",
		TitleCode: "b_chaves",
	}
}
