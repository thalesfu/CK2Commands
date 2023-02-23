package delta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希尔宾ShirbinBarony struct {
	feud.BaseBarony
}

var BShirbin希尔宾 feud.Barony = &希尔宾ShirbinBarony{}

func init() {
	f := BShirbin希尔宾.(*希尔宾ShirbinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shirbin",
		TitleName: "希尔宾",
		TitleCode: "b_shirbin",
	}
}
