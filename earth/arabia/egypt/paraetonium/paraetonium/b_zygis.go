package paraetonium

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卒吉斯ZygisBarony struct {
	feud.BaseBarony
}

var BZygis卒吉斯 feud.Barony = &卒吉斯ZygisBarony{}

func init() {
    f := BZygis卒吉斯.(*卒吉斯ZygisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zygis",
		TitleName: "卒吉斯",
		TitleCode: "b_zygis",
	}
}
