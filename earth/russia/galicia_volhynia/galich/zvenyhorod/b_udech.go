package zvenyhorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌杰奇UdechBarony struct {
	feud.BaseBarony
}

var BUdech乌杰奇 feud.Barony = &乌杰奇UdechBarony{}

func init() {
    f := BUdech乌杰奇.(*乌杰奇UdechBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "udech",
		TitleName: "乌杰奇",
		TitleCode: "b_udech",
	}
}
