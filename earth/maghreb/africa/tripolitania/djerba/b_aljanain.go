package djerba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰纳因AljanainBarony struct {
	feud.BaseBarony
}

var BAljanain杰纳因 feud.Barony = &杰纳因AljanainBarony{}

func init() {
    f := BAljanain杰纳因.(*杰纳因AljanainBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aljanain",
		TitleName: "杰纳因",
		TitleCode: "b_aljanain",
	}
}
