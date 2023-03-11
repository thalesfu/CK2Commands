package il-khanate

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Il-khanateEmpire interface {
    feud.Empire
}

type 伊儿汗国Il-khanateEmpire struct {
	feud.BaseEmpire
}

var EIl-khanate伊儿汗国 Il-khanateEmpire = &伊儿汗国Il-khanateEmpire{}

func init() {
	f := EIl-khanate伊儿汗国.(*伊儿汗国Il-khanateEmpire)
	f.BaseEmpire = feud.BaseEmpire{
		Title:     "il-khanate",
		TitleName: "伊儿汗国",
		TitleCode: "e_il-khanate",
		Kingdoms:  map[string]feud.Kingdom{},
	}
}
