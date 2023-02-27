package sevilla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓜代拉堡AlcaladeguadairaBarony struct {
	feud.BaseBarony
}

var BAlcaladeguadaira瓜代拉堡 feud.Barony = &瓜代拉堡AlcaladeguadairaBarony{}

func init() {
    f := BAlcaladeguadaira瓜代拉堡.(*瓜代拉堡AlcaladeguadairaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alcaladeguadaira",
		TitleName: "瓜代拉堡",
		TitleCode: "b_alcaladeguadaira",
	}
}
