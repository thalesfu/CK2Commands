package tsaparang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托林贡巴TholinggompaBarony struct {
	feud.BaseBarony
}

var BTholinggompa托林贡巴 feud.Barony = &托林贡巴TholinggompaBarony{}

func init() {
    f := BTholinggompa托林贡巴.(*托林贡巴TholinggompaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tholinggompa",
		TitleName: "托林贡巴",
		TitleCode: "b_tholinggompa",
	}
}
