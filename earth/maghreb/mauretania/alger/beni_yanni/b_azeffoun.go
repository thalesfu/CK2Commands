package beni_yanni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿宰丰AzeffounBarony struct {
	feud.BaseBarony
}

var BAzeffoun阿宰丰 feud.Barony = &阿宰丰AzeffounBarony{}

func init() {
    f := BAzeffoun阿宰丰.(*阿宰丰AzeffounBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "azeffoun",
		TitleName: "阿宰丰",
		TitleCode: "b_azeffoun",
	}
}
