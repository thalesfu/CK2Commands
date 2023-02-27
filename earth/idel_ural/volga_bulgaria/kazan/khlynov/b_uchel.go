package khlynov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌切尔UchelBarony struct {
	feud.BaseBarony
}

var BUchel乌切尔 feud.Barony = &乌切尔UchelBarony{}

func init() {
    f := BUchel乌切尔.(*乌切尔UchelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uchel",
		TitleName: "乌切尔",
		TitleCode: "b_uchel",
	}
}
