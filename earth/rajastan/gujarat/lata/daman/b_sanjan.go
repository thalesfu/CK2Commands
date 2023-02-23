package daman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳加尔SanjanBarony struct {
	feud.BaseBarony
}

var BSanjan纳加尔 feud.Barony = &纳加尔SanjanBarony{}

func init() {
	f := BSanjan纳加尔.(*纳加尔SanjanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanjan",
		TitleName: "纳加尔",
		TitleCode: "b_sanjan",
	}
}
