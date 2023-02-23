package reni

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 悉陀目佉SidhmukhBarony struct {
	feud.BaseBarony
}

var BSidhmukh悉陀目佉 feud.Barony = &悉陀目佉SidhmukhBarony{}

func init() {
	f := BSidhmukh悉陀目佉.(*悉陀目佉SidhmukhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidhmukh",
		TitleName: "悉陀目佉",
		TitleCode: "b_sidhmukh",
	}
}
