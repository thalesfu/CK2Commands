package dithmarschen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊策霍ItzehoeBarony struct {
	feud.BaseBarony
}

var BItzehoe伊策霍 feud.Barony = &伊策霍ItzehoeBarony{}

func init() {
    f := BItzehoe伊策霍.(*伊策霍ItzehoeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "itzehoe",
		TitleName: "伊策霍",
		TitleCode: "b_itzehoe",
	}
}
