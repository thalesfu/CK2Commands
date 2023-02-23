package lhatok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 则巴ZebaBarony struct {
	feud.BaseBarony
}

var BZeba则巴 feud.Barony = &则巴ZebaBarony{}

func init() {
	f := BZeba则巴.(*则巴ZebaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zeba",
		TitleName: "则巴",
		TitleCode: "b_zeba",
	}
}
