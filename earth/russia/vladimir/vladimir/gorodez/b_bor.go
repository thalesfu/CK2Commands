package gorodez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博尔BorBarony struct {
	feud.BaseBarony
}

var BBor博尔 feud.Barony = &博尔BorBarony{}

func init() {
    f := BBor博尔.(*博尔BorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bor",
		TitleName: "博尔",
		TitleCode: "b_bor",
	}
}
