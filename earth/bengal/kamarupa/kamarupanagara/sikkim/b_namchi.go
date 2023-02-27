package sikkim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳姆吉NamchiBarony struct {
	feud.BaseBarony
}

var BNamchi纳姆吉 feud.Barony = &纳姆吉NamchiBarony{}

func init() {
    f := BNamchi纳姆吉.(*纳姆吉NamchiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "namchi",
		TitleName: "纳姆吉",
		TitleCode: "b_namchi",
	}
}
