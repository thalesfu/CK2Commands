package zaysan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊万诺夫卡IvanovkaBarony struct {
	feud.BaseBarony
}

var BIvanovka伊万诺夫卡 feud.Barony = &伊万诺夫卡IvanovkaBarony{}

func init() {
	f := BIvanovka伊万诺夫卡.(*伊万诺夫卡IvanovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ivanovka",
		TitleName: "伊万诺夫卡",
		TitleCode: "b_ivanovka",
	}
}
