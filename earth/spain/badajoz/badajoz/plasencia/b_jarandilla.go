package plasencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈兰迪利亚JarandillaBarony struct {
	feud.BaseBarony
}

var BJarandilla哈兰迪利亚 feud.Barony = &哈兰迪利亚JarandillaBarony{}

func init() {
    f := BJarandilla哈兰迪利亚.(*哈兰迪利亚JarandillaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jarandilla",
		TitleName: "哈兰迪利亚",
		TitleCode: "b_jarandilla",
	}
}
