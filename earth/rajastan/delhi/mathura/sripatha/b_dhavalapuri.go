package sripatha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀伐罗补梨DhavalapuriBarony struct {
	feud.BaseBarony
}

var BDhavalapuri陀伐罗补梨 feud.Barony = &陀伐罗补梨DhavalapuriBarony{}

func init() {
	f := BDhavalapuri陀伐罗补梨.(*陀伐罗补梨DhavalapuriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhavalapuri",
		TitleName: "陀伐罗补梨",
		TitleCode: "b_dhavalapuri",
	}
}
