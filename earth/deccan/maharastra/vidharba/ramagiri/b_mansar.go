package ramagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼瑟尔MansarBarony struct {
	feud.BaseBarony
}

var BMansar曼瑟尔 feud.Barony = &曼瑟尔MansarBarony{}

func init() {
	f := BMansar曼瑟尔.(*曼瑟尔MansarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mansar",
		TitleName: "曼瑟尔",
		TitleCode: "b_mansar",
	}
}
