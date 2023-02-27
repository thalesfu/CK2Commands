package salamanca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特拉迪略斯TerradillosBarony struct {
	feud.BaseBarony
}

var BTerradillos特拉迪略斯 feud.Barony = &特拉迪略斯TerradillosBarony{}

func init() {
    f := BTerradillos特拉迪略斯.(*特拉迪略斯TerradillosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "terradillos",
		TitleName: "特拉迪略斯",
		TitleCode: "b_terradillos",
	}
}
