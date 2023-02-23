package shetland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎宁斯堡CunningsburghBarony struct {
	feud.BaseBarony
}

var BCunningsburgh坎宁斯堡 feud.Barony = &坎宁斯堡CunningsburghBarony{}

func init() {
	f := BCunningsburgh坎宁斯堡.(*坎宁斯堡CunningsburghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cunningsburgh",
		TitleName: "坎宁斯堡",
		TitleCode: "b_cunningsburgh",
	}
}
