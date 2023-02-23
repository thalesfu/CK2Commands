package esztergom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 考考特KakathBarony struct {
	feud.BaseBarony
}

var BKakath考考特 feud.Barony = &考考特KakathBarony{}

func init() {
	f := BKakath考考特.(*考考特KakathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kakath",
		TitleName: "考考特",
		TitleCode: "b_kakath",
	}
}
