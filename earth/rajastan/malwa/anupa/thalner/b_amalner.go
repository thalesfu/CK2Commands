package thalner

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿默尔内尔AmalnerBarony struct {
	feud.BaseBarony
}

var BAmalner阿默尔内尔 feud.Barony = &阿默尔内尔AmalnerBarony{}

func init() {
	f := BAmalner阿默尔内尔.(*阿默尔内尔AmalnerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amalner",
		TitleName: "阿默尔内尔",
		TitleCode: "b_amalner",
	}
}
