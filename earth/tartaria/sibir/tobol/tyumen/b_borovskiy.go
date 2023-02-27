package tyumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博罗夫斯基BorovskiyBarony struct {
	feud.BaseBarony
}

var BBorovskiy博罗夫斯基 feud.Barony = &博罗夫斯基BorovskiyBarony{}

func init() {
    f := BBorovskiy博罗夫斯基.(*博罗夫斯基BorovskiyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "borovskiy",
		TitleName: "博罗夫斯基",
		TitleCode: "b_borovskiy",
	}
}
