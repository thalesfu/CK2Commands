package karur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维哈里VehariBarony struct {
	feud.BaseBarony
}

var BVehari维哈里 feud.Barony = &维哈里VehariBarony{}

func init() {
    f := BVehari维哈里.(*维哈里VehariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vehari",
		TitleName: "维哈里",
		TitleCode: "b_vehari",
	}
}
