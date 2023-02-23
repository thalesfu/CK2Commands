package jeddah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉达JeddahBarony struct {
	feud.BaseBarony
}

var BJeddah吉达 feud.Barony = &吉达JeddahBarony{}

func init() {
	f := BJeddah吉达.(*吉达JeddahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jeddah",
		TitleName: "吉达",
		TitleCode: "b_jeddah",
	}
}
