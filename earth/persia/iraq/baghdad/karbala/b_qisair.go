package karbala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇瑟QisairBarony struct {
	feud.BaseBarony
}

var BQisair奇瑟 feud.Barony = &奇瑟QisairBarony{}

func init() {
	f := BQisair奇瑟.(*奇瑟QisairBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qisair",
		TitleName: "奇瑟",
		TitleCode: "b_qisair",
	}
}
