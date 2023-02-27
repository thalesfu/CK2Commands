package lumbini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 都罗私补罗TulsipurBarony struct {
	feud.BaseBarony
}

var BTulsipur都罗私补罗 feud.Barony = &都罗私补罗TulsipurBarony{}

func init() {
    f := BTulsipur都罗私补罗.(*都罗私补罗TulsipurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tulsipur",
		TitleName: "都罗私补罗",
		TitleCode: "b_tulsipur",
	}
}
