package barasuru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼俱侯支ManghechBarony struct {
	feud.BaseBarony
}

var BManghech曼俱侯支 feud.Barony = &曼俱侯支ManghechBarony{}

func init() {
    f := BManghech曼俱侯支.(*曼俱侯支ManghechBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manghech",
		TitleName: "曼俱侯支",
		TitleCode: "b_manghech",
	}
}
