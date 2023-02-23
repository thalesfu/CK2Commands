package calatrava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 麦德林MedellinBarony struct {
	feud.BaseBarony
}

var BMedellin麦德林 feud.Barony = &麦德林MedellinBarony{}

func init() {
	f := BMedellin麦德林.(*麦德林MedellinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medellin",
		TitleName: "麦德林",
		TitleCode: "b_medellin",
	}
}
