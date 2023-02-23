package ranthambore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗那萨担婆补罗RanthamboreBarony struct {
	feud.BaseBarony
}

var BRanthambore罗那萨担婆补罗 feud.Barony = &罗那萨担婆补罗RanthamboreBarony{}

func init() {
	f := BRanthambore罗那萨担婆补罗.(*罗那萨担婆补罗RanthamboreBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ranthambore",
		TitleName: "罗那萨担婆补罗",
		TitleCode: "b_ranthambore",
	}
}
