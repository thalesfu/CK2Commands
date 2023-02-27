package pundravardhana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏摩补罗摩诃毗诃罗Somapura_mahaviharaBarony struct {
	feud.BaseBarony
}

var BSomapura_mahavihara苏摩补罗摩诃毗诃罗 feud.Barony = &苏摩补罗摩诃毗诃罗Somapura_mahaviharaBarony{}

func init() {
    f := BSomapura_mahavihara苏摩补罗摩诃毗诃罗.(*苏摩补罗摩诃毗诃罗Somapura_mahaviharaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "somapura_mahavihara",
		TitleName: "苏摩补罗摩诃毗诃罗",
		TitleCode: "b_somapura_mahavihara",
	}
}
