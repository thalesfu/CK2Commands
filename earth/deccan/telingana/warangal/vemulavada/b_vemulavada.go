package vemulavada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠牟罗婆荼VemulavadaBarony struct {
	feud.BaseBarony
}

var BVemulavada吠牟罗婆荼 feud.Barony = &吠牟罗婆荼VemulavadaBarony{}

func init() {
    f := BVemulavada吠牟罗婆荼.(*吠牟罗婆荼VemulavadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vemulavada",
		TitleName: "吠牟罗婆荼",
		TitleCode: "b_vemulavada",
	}
}
