package mahoyadapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤陀伽UdagaiBarony struct {
	feud.BaseBarony
}

var BUdagai尤陀伽 feud.Barony = &尤陀伽UdagaiBarony{}

func init() {
    f := BUdagai尤陀伽.(*尤陀伽UdagaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "udagai",
		TitleName: "尤陀伽",
		TitleCode: "b_udagai",
	}
}
