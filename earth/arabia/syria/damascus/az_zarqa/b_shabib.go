package az_zarqa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙比布ShabibBarony struct {
	feud.BaseBarony
}

var BShabib沙比布 feud.Barony = &沙比布ShabibBarony{}

func init() {
    f := BShabib沙比布.(*沙比布ShabibBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shabib",
		TitleName: "沙比布",
		TitleCode: "b_shabib",
	}
}
