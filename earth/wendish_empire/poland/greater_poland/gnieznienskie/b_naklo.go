package gnieznienskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳克沃NakloBarony struct {
	feud.BaseBarony
}

var BNaklo纳克沃 feud.Barony = &纳克沃NakloBarony{}

func init() {
    f := BNaklo纳克沃.(*纳克沃NakloBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naklo",
		TitleName: "纳克沃",
		TitleCode: "b_naklo",
	}
}
