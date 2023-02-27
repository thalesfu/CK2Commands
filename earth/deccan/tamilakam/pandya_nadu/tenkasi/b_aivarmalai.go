package tenkasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伐摩罗AivarmalaiBarony struct {
	feud.BaseBarony
}

var BAivarmalai阿伐摩罗 feud.Barony = &阿伐摩罗AivarmalaiBarony{}

func init() {
    f := BAivarmalai阿伐摩罗.(*阿伐摩罗AivarmalaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aivarmalai",
		TitleName: "阿伐摩罗",
		TitleCode: "b_aivarmalai",
	}
}
