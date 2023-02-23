package seres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫里斯图波利斯ChristoupolisBarony struct {
	feud.BaseBarony
}

var BChristoupolis赫里斯图波利斯 feud.Barony = &赫里斯图波利斯ChristoupolisBarony{}

func init() {
	f := BChristoupolis赫里斯图波利斯.(*赫里斯图波利斯ChristoupolisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "christoupolis",
		TitleName: "赫里斯图波利斯",
		TitleCode: "b_christoupolis",
	}
}
