package olomouc

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥洛穆茨OlomoucBarony struct {
	feud.BaseBarony
}

var BOlomouc奥洛穆茨 feud.Barony = &奥洛穆茨OlomoucBarony{}

func init() {
    f := BOlomouc奥洛穆茨.(*奥洛穆茨OlomoucBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "olomouc",
		TitleName: "奥洛穆茨",
		TitleCode: "b_olomouc",
	}
}
