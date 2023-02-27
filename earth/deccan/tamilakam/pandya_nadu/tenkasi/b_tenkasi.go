package tenkasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谛那迦尸TenkasiBarony struct {
	feud.BaseBarony
}

var BTenkasi谛那迦尸 feud.Barony = &谛那迦尸TenkasiBarony{}

func init() {
    f := BTenkasi谛那迦尸.(*谛那迦尸TenkasiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tenkasi",
		TitleName: "谛那迦尸",
		TitleCode: "b_tenkasi",
	}
}
