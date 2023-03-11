package il_khanate

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Il_khanateEmpire interface {
    feud.Empire
}

type 伊儿汗国Il_khanateEmpire struct {
	feud.BaseEmpire
}

var EIl_khanate伊儿汗国 Il_khanateEmpire = &伊儿汗国Il_khanateEmpire{}

func init() {
	f := EIl_khanate伊儿汗国.(*伊儿汗国Il_khanateEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "il_khanate",
		TitleName: "伊儿汗国",
		TitleCode: "e_il-khanate",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
