package guria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌达比诺UdabnoBarony struct {
	feud.BaseBarony
}

var BUdabno乌达比诺 feud.Barony = &乌达比诺UdabnoBarony{}

func init() {
    f := BUdabno乌达比诺.(*乌达比诺UdabnoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "udabno",
		TitleName: "乌达比诺",
		TitleCode: "b_udabno",
	}
}
