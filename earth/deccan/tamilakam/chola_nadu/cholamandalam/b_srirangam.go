package cholamandalam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯里愣甘SrirangamBarony struct {
	feud.BaseBarony
}

var BSrirangam斯里愣甘 feud.Barony = &斯里愣甘SrirangamBarony{}

func init() {
    f := BSrirangam斯里愣甘.(*斯里愣甘SrirangamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "srirangam",
		TitleName: "斯里愣甘",
		TitleCode: "b_srirangam",
	}
}
