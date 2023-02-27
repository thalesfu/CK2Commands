package kathiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希赫尔ShiharBarony struct {
	feud.BaseBarony
}

var BShihar希赫尔 feud.Barony = &希赫尔ShiharBarony{}

func init() {
    f := BShihar希赫尔.(*希赫尔ShiharBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shihar",
		TitleName: "希赫尔",
		TitleCode: "b_shihar",
	}
}
