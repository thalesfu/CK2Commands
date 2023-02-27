package chilia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利斯基LiskyBarony struct {
	feud.BaseBarony
}

var BLisky利斯基 feud.Barony = &利斯基LiskyBarony{}

func init() {
    f := BLisky利斯基.(*利斯基LiskyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lisky",
		TitleName: "利斯基",
		TitleCode: "b_lisky",
	}
}
