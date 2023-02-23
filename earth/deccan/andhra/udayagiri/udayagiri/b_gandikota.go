package udayagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 甘地科塔GandikotaBarony struct {
	feud.BaseBarony
}

var BGandikota甘地科塔 feud.Barony = &甘地科塔GandikotaBarony{}

func init() {
	f := BGandikota甘地科塔.(*甘地科塔GandikotaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gandikota",
		TitleName: "甘地科塔",
		TitleCode: "b_gandikota",
	}
}
