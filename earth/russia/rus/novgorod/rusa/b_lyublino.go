package rusa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柳布利诺LyublinoBarony struct {
	feud.BaseBarony
}

var BLyublino柳布利诺 feud.Barony = &柳布利诺LyublinoBarony{}

func init() {
	f := BLyublino柳布利诺.(*柳布利诺LyublinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyublino",
		TitleName: "柳布利诺",
		TitleCode: "b_lyublino",
	}
}
