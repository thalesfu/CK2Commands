package messina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特罗伊纳TroiniaBarony struct {
	feud.BaseBarony
}

var BTroinia特罗伊纳 feud.Barony = &特罗伊纳TroiniaBarony{}

func init() {
    f := BTroinia特罗伊纳.(*特罗伊纳TroiniaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "troinia",
		TitleName: "特罗伊纳",
		TitleCode: "b_troinia",
	}
}
