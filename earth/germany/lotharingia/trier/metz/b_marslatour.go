package metz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马尔拉图MarslatourBarony struct {
	feud.BaseBarony
}

var BMarslatour马尔拉图 feud.Barony = &马尔拉图MarslatourBarony{}

func init() {
	f := BMarslatour马尔拉图.(*马尔拉图MarslatourBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marslatour",
		TitleName: "马尔拉图",
		TitleCode: "b_marslatour",
	}
}
