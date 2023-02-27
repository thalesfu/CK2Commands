package naimisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆诃罗伊支BahraichBarony struct {
	feud.BaseBarony
}

var BBahraich婆诃罗伊支 feud.Barony = &婆诃罗伊支BahraichBarony{}

func init() {
    f := BBahraich婆诃罗伊支.(*婆诃罗伊支BahraichBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bahraich",
		TitleName: "婆诃罗伊支",
		TitleCode: "b_bahraich",
	}
}
