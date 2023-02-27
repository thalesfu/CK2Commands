package plzen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯特日布罗StriboBarony struct {
	feud.BaseBarony
}

var BStribo斯特日布罗 feud.Barony = &斯特日布罗StriboBarony{}

func init() {
    f := BStribo斯特日布罗.(*斯特日布罗StriboBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stribo",
		TitleName: "斯特日布罗",
		TitleCode: "b_stribo",
	}
}
