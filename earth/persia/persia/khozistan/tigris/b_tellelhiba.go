package tigris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特尔赫巴TellelhibaBarony struct {
	feud.BaseBarony
}

var BTellelhiba特尔赫巴 feud.Barony = &特尔赫巴TellelhibaBarony{}

func init() {
    f := BTellelhiba特尔赫巴.(*特尔赫巴TellelhibaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tellelhiba",
		TitleName: "特尔赫巴",
		TitleCode: "b_tellelhiba",
	}
}
