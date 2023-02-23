package mozhaysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 通沙耶沃TonshaevoBarony struct {
	feud.BaseBarony
}

var BTonshaevo通沙耶沃 feud.Barony = &通沙耶沃TonshaevoBarony{}

func init() {
	f := BTonshaevo通沙耶沃.(*通沙耶沃TonshaevoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tonshaevo",
		TitleName: "通沙耶沃",
		TitleCode: "b_tonshaevo",
	}
}
