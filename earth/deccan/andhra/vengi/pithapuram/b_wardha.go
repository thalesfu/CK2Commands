package pithapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 跋罗陀WardhaBarony struct {
	feud.BaseBarony
}

var BWardha跋罗陀 feud.Barony = &跋罗陀WardhaBarony{}

func init() {
	f := BWardha跋罗陀.(*跋罗陀WardhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wardha",
		TitleName: "跋罗陀",
		TitleCode: "b_wardha",
	}
}
