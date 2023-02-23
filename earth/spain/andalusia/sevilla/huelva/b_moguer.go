package huelva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫格尔MoguerBarony struct {
	feud.BaseBarony
}

var BMoguer莫格尔 feud.Barony = &莫格尔MoguerBarony{}

func init() {
	f := BMoguer莫格尔.(*莫格尔MoguerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "moguer",
		TitleName: "莫格尔",
		TitleCode: "b_moguer",
	}
}
