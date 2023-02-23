package buhairya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊布沙韦IbshawaiBarony struct {
	feud.BaseBarony
}

var BIbshawai伊布沙韦 feud.Barony = &伊布沙韦IbshawaiBarony{}

func init() {
	f := BIbshawai伊布沙韦.(*伊布沙韦IbshawaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ibshawai",
		TitleName: "伊布沙韦",
		TitleCode: "b_ibshawai",
	}
}
