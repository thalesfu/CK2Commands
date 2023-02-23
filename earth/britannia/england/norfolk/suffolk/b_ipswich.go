package suffolk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊普斯威奇IpswichBarony struct {
	feud.BaseBarony
}

var BIpswich伊普斯威奇 feud.Barony = &伊普斯威奇IpswichBarony{}

func init() {
	f := BIpswich伊普斯威奇.(*伊普斯威奇IpswichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ipswich",
		TitleName: "伊普斯威奇",
		TitleCode: "b_ipswich",
	}
}
