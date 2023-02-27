package katsina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富拉纳FuranaBarony struct {
	feud.BaseBarony
}

var BFurana富拉纳 feud.Barony = &富拉纳FuranaBarony{}

func init() {
    f := BFurana富拉纳.(*富拉纳FuranaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "furana",
		TitleName: "富拉纳",
		TitleCode: "b_furana",
	}
}
