package kongu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠卢尔PerurBarony struct {
	feud.BaseBarony
}

var BPerur吠卢尔 feud.Barony = &吠卢尔PerurBarony{}

func init() {
    f := BPerur吠卢尔.(*吠卢尔PerurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perur",
		TitleName: "吠卢尔",
		TitleCode: "b_perur",
	}
}
