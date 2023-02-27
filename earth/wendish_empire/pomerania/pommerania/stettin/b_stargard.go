package stettin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施塔加德StargardBarony struct {
	feud.BaseBarony
}

var BStargard施塔加德 feud.Barony = &施塔加德StargardBarony{}

func init() {
    f := BStargard施塔加德.(*施塔加德StargardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stargard",
		TitleName: "施塔加德",
		TitleCode: "b_stargard",
	}
}
