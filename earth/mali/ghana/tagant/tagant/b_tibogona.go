package tagant

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂博戈纳TibogonaBarony struct {
	feud.BaseBarony
}

var BTibogona蒂博戈纳 feud.Barony = &蒂博戈纳TibogonaBarony{}

func init() {
    f := BTibogona蒂博戈纳.(*蒂博戈纳TibogonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tibogona",
		TitleName: "蒂博戈纳",
		TitleCode: "b_tibogona",
	}
}
