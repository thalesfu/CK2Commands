package liege

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那慕尔NamurBarony struct {
	feud.BaseBarony
}

var BNamur那慕尔 feud.Barony = &那慕尔NamurBarony{}

func init() {
    f := BNamur那慕尔.(*那慕尔NamurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "namur",
		TitleName: "那慕尔",
		TitleCode: "b_namur",
	}
}
