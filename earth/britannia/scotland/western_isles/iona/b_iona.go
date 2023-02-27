package iona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾奥纳IonaBarony struct {
	feud.BaseBarony
}

var BIona艾奥纳 feud.Barony = &艾奥纳IonaBarony{}

func init() {
    f := BIona艾奥纳.(*艾奥纳IonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iona",
		TitleName: "艾奥纳",
		TitleCode: "b_iona",
	}
}
