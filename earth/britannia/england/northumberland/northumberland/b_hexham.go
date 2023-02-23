package northumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 赫克瑟姆HexhamBarony struct {
	feud.BaseBarony
}

var BHexham赫克瑟姆 feud.Barony = &赫克瑟姆HexhamBarony{}

func init() {
	f := BHexham赫克瑟姆.(*赫克瑟姆HexhamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hexham",
		TitleName: "赫克瑟姆",
		TitleCode: "b_hexham",
	}
}
