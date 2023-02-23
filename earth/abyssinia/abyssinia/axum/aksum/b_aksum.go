package aksum

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿克苏姆AksumBarony struct {
	feud.BaseBarony
}

var BAksum阿克苏姆 feud.Barony = &阿克苏姆AksumBarony{}

func init() {
	f := BAksum阿克苏姆.(*阿克苏姆AksumBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aksum",
		TitleName: "阿克苏姆",
		TitleCode: "b_aksum",
	}
}
