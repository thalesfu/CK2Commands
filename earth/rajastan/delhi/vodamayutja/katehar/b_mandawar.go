package katehar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼荼婆罗MandawarBarony struct {
	feud.BaseBarony
}

var BMandawar曼荼婆罗 feud.Barony = &曼荼婆罗MandawarBarony{}

func init() {
	f := BMandawar曼荼婆罗.(*曼荼婆罗MandawarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mandawar",
		TitleName: "曼荼婆罗",
		TitleCode: "b_mandawar",
	}
}
