package ratanpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 翳兰头YellanduBarony struct {
	feud.BaseBarony
}

var BYellandu翳兰头 feud.Barony = &翳兰头YellanduBarony{}

func init() {
	f := BYellandu翳兰头.(*翳兰头YellanduBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yellandu",
		TitleName: "翳兰头",
		TitleCode: "b_yellandu",
	}
}
