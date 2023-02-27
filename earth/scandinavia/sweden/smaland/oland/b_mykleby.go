package oland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米克雷比MyklebyBarony struct {
	feud.BaseBarony
}

var BMykleby米克雷比 feud.Barony = &米克雷比MyklebyBarony{}

func init() {
    f := BMykleby米克雷比.(*米克雷比MyklebyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mykleby",
		TitleName: "米克雷比",
		TitleCode: "b_mykleby",
	}
}
