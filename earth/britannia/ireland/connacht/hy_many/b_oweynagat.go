package hy_many

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌埃夫纳加特OweynagatBarony struct {
	feud.BaseBarony
}

var BOweynagat乌埃夫纳加特 feud.Barony = &乌埃夫纳加特OweynagatBarony{}

func init() {
    f := BOweynagat乌埃夫纳加特.(*乌埃夫纳加特OweynagatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "oweynagat",
		TitleName: "乌埃夫纳加特",
		TitleCode: "b_oweynagat",
	}
}
