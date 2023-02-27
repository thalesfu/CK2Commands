package mahoyadapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伽毗由KaviyurBarony struct {
	feud.BaseBarony
}

var BKaviyur伽毗由 feud.Barony = &伽毗由KaviyurBarony{}

func init() {
    f := BKaviyur伽毗由.(*伽毗由KaviyurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kaviyur",
		TitleName: "伽毗由",
		TitleCode: "b_kaviyur",
	}
}
