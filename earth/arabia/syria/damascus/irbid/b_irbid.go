package irbid

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊尔比德IrbidBarony struct {
	feud.BaseBarony
}

var BIrbid伊尔比德 feud.Barony = &伊尔比德IrbidBarony{}

func init() {
    f := BIrbid伊尔比德.(*伊尔比德IrbidBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "irbid",
		TitleName: "伊尔比德",
		TitleCode: "b_irbid",
	}
}
