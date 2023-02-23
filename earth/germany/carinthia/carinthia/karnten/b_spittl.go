package karnten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯皮特尔SpittlBarony struct {
	feud.BaseBarony
}

var BSpittl斯皮特尔 feud.Barony = &斯皮特尔SpittlBarony{}

func init() {
	f := BSpittl斯皮特尔.(*斯皮特尔SpittlBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "spittl",
		TitleName: "斯皮特尔",
		TitleCode: "b_spittl",
	}
}
