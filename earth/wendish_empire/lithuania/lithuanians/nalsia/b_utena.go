package nalsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌田纳UtenaBarony struct {
	feud.BaseBarony
}

var BUtena乌田纳 feud.Barony = &乌田纳UtenaBarony{}

func init() {
    f := BUtena乌田纳.(*乌田纳UtenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "utena",
		TitleName: "乌田纳",
		TitleCode: "b_utena",
	}
}
