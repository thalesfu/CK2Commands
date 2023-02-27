package siuntio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 万塔VantaaBarony struct {
	feud.BaseBarony
}

var BVantaa万塔 feud.Barony = &万塔VantaaBarony{}

func init() {
    f := BVantaa万塔.(*万塔VantaaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vantaa",
		TitleName: "万塔",
		TitleCode: "b_vantaa",
	}
}
