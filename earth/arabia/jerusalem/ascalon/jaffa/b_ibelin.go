package jaffa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊贝林IbelinBarony struct {
	feud.BaseBarony
}

var BIbelin伊贝林 feud.Barony = &伊贝林IbelinBarony{}

func init() {
	f := BIbelin伊贝林.(*伊贝林IbelinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ibelin",
		TitleName: "伊贝林",
		TitleCode: "b_ibelin",
	}
}
