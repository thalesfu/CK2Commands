package asturias_de_oviedo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳尔塞阿河畔坎加斯CangasdelnarceaBarony struct {
	feud.BaseBarony
}

var BCangasdelnarcea纳尔塞阿河畔坎加斯 feud.Barony = &纳尔塞阿河畔坎加斯CangasdelnarceaBarony{}

func init() {
    f := BCangasdelnarcea纳尔塞阿河畔坎加斯.(*纳尔塞阿河畔坎加斯CangasdelnarceaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cangasdelnarcea",
		TitleName: "纳尔塞阿河畔坎加斯",
		TitleCode: "b_cangasdelnarcea",
	}
}
