package meknes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿穆盖尔AmouguerBarony struct {
	feud.BaseBarony
}

var BAmouguer阿穆盖尔 feud.Barony = &阿穆盖尔AmouguerBarony{}

func init() {
	f := BAmouguer阿穆盖尔.(*阿穆盖尔AmouguerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amouguer",
		TitleName: "阿穆盖尔",
		TitleCode: "b_amouguer",
	}
}
