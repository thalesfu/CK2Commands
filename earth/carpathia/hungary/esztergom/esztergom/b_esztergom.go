package esztergom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃斯泰尔戈姆EsztergomBarony struct {
	feud.BaseBarony
}

var BEsztergom埃斯泰尔戈姆 feud.Barony = &埃斯泰尔戈姆EsztergomBarony{}

func init() {
	f := BEsztergom埃斯泰尔戈姆.(*埃斯泰尔戈姆EsztergomBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "esztergom",
		TitleName: "埃斯泰尔戈姆",
		TitleCode: "b_esztergom",
	}
}
