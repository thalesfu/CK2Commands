package hail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉奇丹IqdahBarony struct {
	feud.BaseBarony
}

var BIqdah拉奇丹 feud.Barony = &拉奇丹IqdahBarony{}

func init() {
	f := BIqdah拉奇丹.(*拉奇丹IqdahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "iqdah",
		TitleName: "拉奇丹",
		TitleCode: "b_iqdah",
	}
}
