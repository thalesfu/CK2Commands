package porto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蓬蒂迪利马PontedelimaBarony struct {
	feud.BaseBarony
}

var BPontedelima蓬蒂迪利马 feud.Barony = &蓬蒂迪利马PontedelimaBarony{}

func init() {
    f := BPontedelima蓬蒂迪利马.(*蓬蒂迪利马PontedelimaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pontedelima",
		TitleName: "蓬蒂迪利马",
		TitleCode: "b_pontedelima",
	}
}
