package lechczechrus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LechczechrusEmpire interface {
    feud.Empire
}

type 斯拉夫联盟LechczechrusEmpire struct {
	feud.BaseEmpire
}

var ELechczechrus斯拉夫联盟 LechczechrusEmpire = &斯拉夫联盟LechczechrusEmpire{}

func init() {
	f := ELechczechrus斯拉夫联盟.(*斯拉夫联盟LechczechrusEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "lechczechrus",
		TitleName: "斯拉夫联盟",
		TitleCode: "e_lechczechrus",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
