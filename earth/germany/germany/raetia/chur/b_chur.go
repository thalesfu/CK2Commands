package chur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库尔ChurBarony struct {
	feud.BaseBarony
}

var BChur库尔 feud.Barony = &库尔ChurBarony{}

func init() {
    f := BChur库尔.(*库尔ChurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chur",
		TitleName: "库尔",
		TitleCode: "b_chur",
	}
}
