package taktse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曲沟QoiBarony struct {
	feud.BaseBarony
}

var BQoi曲沟 feud.Barony = &曲沟QoiBarony{}

func init() {
	f := BQoi曲沟.(*曲沟QoiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qoi",
		TitleName: "曲沟",
		TitleCode: "b_qoi",
	}
}
