package yopurga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽师JiashiBarony struct {
	feud.BaseBarony
}

var BJiashi伽师 feud.Barony = &伽师JiashiBarony{}

func init() {
	f := BJiashi伽师.(*伽师JiashiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jiashi",
		TitleName: "伽师",
		TitleCode: "b_jiashi",
	}
}
