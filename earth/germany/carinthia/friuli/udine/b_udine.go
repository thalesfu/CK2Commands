package udine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌迪内UdineBarony struct {
	feud.BaseBarony
}

var BUdine乌迪内 feud.Barony = &乌迪内UdineBarony{}

func init() {
	f := BUdine乌迪内.(*乌迪内UdineBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "udine",
		TitleName: "乌迪内",
		TitleCode: "b_udine",
	}
}
