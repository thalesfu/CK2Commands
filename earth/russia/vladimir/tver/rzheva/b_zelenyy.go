package rzheva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽廖内ZelenyyBarony struct {
	feud.BaseBarony
}

var BZelenyy泽廖内 feud.Barony = &泽廖内ZelenyyBarony{}

func init() {
    f := BZelenyy泽廖内.(*泽廖内ZelenyyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zelenyy",
		TitleName: "泽廖内",
		TitleCode: "b_zelenyy",
	}
}
