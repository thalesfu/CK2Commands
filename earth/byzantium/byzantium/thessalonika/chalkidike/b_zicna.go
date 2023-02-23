package chalkidike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基克纳ZicnaBarony struct {
	feud.BaseBarony
}

var BZicna基克纳 feud.Barony = &基克纳ZicnaBarony{}

func init() {
	f := BZicna基克纳.(*基克纳ZicnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zicna",
		TitleName: "基克纳",
		TitleCode: "b_zicna",
	}
}
