package mainz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MainzDuke interface {
    feud.Duke
}

type 美因茨MainzDuke struct {
	feud.BaseDuke
}

var DMainz美因茨 MainzDuke = &美因茨MainzDuke{}

func init() {
	f := DMainz美因茨.(*美因茨MainzDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "mainz",
		TitleName: "美因茨",
		TitleCode: "d_mainz",
		Counties:  map[string]feud.County{},
	}

}
