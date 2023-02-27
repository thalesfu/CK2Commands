package kano

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡诺KanoBarony struct {
	feud.BaseBarony
}

var BKano卡诺 feud.Barony = &卡诺KanoBarony{}

func init() {
    f := BKano卡诺.(*卡诺KanoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kano",
		TitleName: "卡诺",
		TitleCode: "b_kano",
	}
}
