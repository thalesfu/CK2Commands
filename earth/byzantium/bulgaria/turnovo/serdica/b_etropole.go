package serdica

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃特罗波莱EtropoleBarony struct {
	feud.BaseBarony
}

var BEtropole埃特罗波莱 feud.Barony = &埃特罗波莱EtropoleBarony{}

func init() {
	f := BEtropole埃特罗波莱.(*埃特罗波莱EtropoleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "etropole",
		TitleName: "埃特罗波莱",
		TitleCode: "b_etropole",
	}
}
