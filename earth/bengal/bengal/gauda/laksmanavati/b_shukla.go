package laksmanavati

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 叔罗ShuklaBarony struct {
	feud.BaseBarony
}

var BShukla叔罗 feud.Barony = &叔罗ShuklaBarony{}

func init() {
    f := BShukla叔罗.(*叔罗ShuklaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shukla",
		TitleName: "叔罗",
		TitleCode: "b_shukla",
	}
}
