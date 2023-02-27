package gabes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉罕AlqalahBarony struct {
	feud.BaseBarony
}

var BAlqalah卡拉罕 feud.Barony = &卡拉罕AlqalahBarony{}

func init() {
    f := BAlqalah卡拉罕.(*卡拉罕AlqalahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alqalah",
		TitleName: "卡拉罕",
		TitleCode: "b_alqalah",
	}
}
