package izborsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍赫雷KhokhlyBarony struct {
	feud.BaseBarony
}

var BKhokhly霍赫雷 feud.Barony = &霍赫雷KhokhlyBarony{}

func init() {
    f := BKhokhly霍赫雷.(*霍赫雷KhokhlyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khokhly",
		TitleName: "霍赫雷",
		TitleCode: "b_khokhly",
	}
}
