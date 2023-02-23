package kurdistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杜胡克DuhokBarony struct {
	feud.BaseBarony
}

var BDuhok杜胡克 feud.Barony = &杜胡克DuhokBarony{}

func init() {
	f := BDuhok杜胡克.(*杜胡克DuhokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "duhok",
		TitleName: "杜胡克",
		TitleCode: "b_duhok",
	}
}
