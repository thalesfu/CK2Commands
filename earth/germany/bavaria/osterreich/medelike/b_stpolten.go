package medelike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣珀尔滕StpoltenBarony struct {
	feud.BaseBarony
}

var BStpolten圣珀尔滕 feud.Barony = &圣珀尔滕StpoltenBarony{}

func init() {
	f := BStpolten圣珀尔滕.(*圣珀尔滕StpoltenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stpolten",
		TitleName: "圣珀尔滕",
		TitleCode: "b_stpolten",
	}
}
