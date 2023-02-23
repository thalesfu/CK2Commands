package hajr

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾卜哈AbhaBarony struct {
	feud.BaseBarony
}

var BAbha艾卜哈 feud.Barony = &艾卜哈AbhaBarony{}

func init() {
	f := BAbha艾卜哈.(*艾卜哈AbhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abha",
		TitleName: "艾卜哈",
		TitleCode: "b_abha",
	}
}
