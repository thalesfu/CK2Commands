package lumbini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陪胪诃伐BhairahawaBarony struct {
	feud.BaseBarony
}

var BBhairahawa陪胪诃伐 feud.Barony = &陪胪诃伐BhairahawaBarony{}

func init() {
    f := BBhairahawa陪胪诃伐.(*陪胪诃伐BhairahawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bhairahawa",
		TitleName: "陪胪诃伐",
		TitleCode: "b_bhairahawa",
	}
}
