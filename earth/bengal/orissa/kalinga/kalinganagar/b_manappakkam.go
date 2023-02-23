package kalinganagar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩那波甘ManappakkamBarony struct {
	feud.BaseBarony
}

var BManappakkam摩那波甘 feud.Barony = &摩那波甘ManappakkamBarony{}

func init() {
	f := BManappakkam摩那波甘.(*摩那波甘ManappakkamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manappakkam",
		TitleName: "摩那波甘",
		TitleCode: "b_manappakkam",
	}
}
