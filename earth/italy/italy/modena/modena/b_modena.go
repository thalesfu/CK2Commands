package modena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩德纳ModenaBarony struct {
	feud.BaseBarony
}

var BModena摩德纳 feud.Barony = &摩德纳ModenaBarony{}

func init() {
    f := BModena摩德纳.(*摩德纳ModenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "modena",
		TitleName: "摩德纳",
		TitleCode: "b_modena",
	}
}
