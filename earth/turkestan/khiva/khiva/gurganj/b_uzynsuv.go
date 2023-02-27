package gurganj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌津苏夫UzynsuvBarony struct {
	feud.BaseBarony
}

var BUzynsuv乌津苏夫 feud.Barony = &乌津苏夫UzynsuvBarony{}

func init() {
    f := BUzynsuv乌津苏夫.(*乌津苏夫UzynsuvBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uzynsuv",
		TitleName: "乌津苏夫",
		TitleCode: "b_uzynsuv",
	}
}
