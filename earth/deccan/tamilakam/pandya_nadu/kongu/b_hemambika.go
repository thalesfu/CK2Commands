package kongu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃曼迦毗HemambikaBarony struct {
	feud.BaseBarony
}

var BHemambika诃曼迦毗 feud.Barony = &诃曼迦毗HemambikaBarony{}

func init() {
    f := BHemambika诃曼迦毗.(*诃曼迦毗HemambikaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hemambika",
		TitleName: "诃曼迦毗",
		TitleCode: "b_hemambika",
	}
}
