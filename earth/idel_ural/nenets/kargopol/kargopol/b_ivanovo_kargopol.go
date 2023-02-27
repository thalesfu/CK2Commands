package kargopol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊万诺沃Ivanovo_kargopolBarony struct {
	feud.BaseBarony
}

var BIvanovo_kargopol伊万诺沃 feud.Barony = &伊万诺沃Ivanovo_kargopolBarony{}

func init() {
    f := BIvanovo_kargopol伊万诺沃.(*伊万诺沃Ivanovo_kargopolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ivanovo_kargopol",
		TitleName: "伊万诺沃",
		TitleCode: "b_ivanovo_kargopol",
	}
}
