package durham_palatine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Durham_palatineDuke interface {
    feud.Duke
}

type 达勒姆Durham_palatineDuke struct {
	feud.BaseDuke
}

var DDurham_palatine达勒姆 Durham_palatineDuke = &达勒姆Durham_palatineDuke{}

func init() {
	f := DDurham_palatine达勒姆.(*达勒姆Durham_palatineDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "durham_palatine",
		TitleName: "达勒姆",
		TitleCode: "d_durham_palatine",
		Counties:  map[string]feud.County{},
	}

}
