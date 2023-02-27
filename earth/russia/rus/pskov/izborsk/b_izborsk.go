package izborsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊兹博尔斯克IzborskBarony struct {
	feud.BaseBarony
}

var BIzborsk伊兹博尔斯克 feud.Barony = &伊兹博尔斯克IzborskBarony{}

func init() {
    f := BIzborsk伊兹博尔斯克.(*伊兹博尔斯克IzborskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "izborsk",
		TitleName: "伊兹博尔斯克",
		TitleCode: "b_izborsk",
	}
}
