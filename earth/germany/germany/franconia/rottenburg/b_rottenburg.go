package rottenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗滕堡RottenburgBarony struct {
	feud.BaseBarony
}

var BRottenburg罗滕堡 feud.Barony = &罗滕堡RottenburgBarony{}

func init() {
    f := BRottenburg罗滕堡.(*罗滕堡RottenburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rottenburg",
		TitleName: "罗滕堡",
		TitleCode: "b_rottenburg",
	}
}
