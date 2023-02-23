package onogost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德尔斯尼克DesnikBarony struct {
	feud.BaseBarony
}

var BDesnik德尔斯尼克 feud.Barony = &德尔斯尼克DesnikBarony{}

func init() {
	f := BDesnik德尔斯尼克.(*德尔斯尼克DesnikBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "desnik",
		TitleName: "德尔斯尼克",
		TitleCode: "b_desnik",
	}
}
