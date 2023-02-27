package khovsgol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库苏古尔KhovsgolBarony struct {
	feud.BaseBarony
}

var BKhovsgol库苏古尔 feud.Barony = &库苏古尔KhovsgolBarony{}

func init() {
    f := BKhovsgol库苏古尔.(*库苏古尔KhovsgolBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khovsgol",
		TitleName: "库苏古尔",
		TitleCode: "b_khovsgol",
	}
}
