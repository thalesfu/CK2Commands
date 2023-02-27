package hajar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊卜里IbriBarony struct {
	feud.BaseBarony
}

var BIbri伊卜里 feud.Barony = &伊卜里IbriBarony{}

func init() {
    f := BIbri伊卜里.(*伊卜里IbriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ibri",
		TitleName: "伊卜里",
		TitleCode: "b_ibri",
	}
}
