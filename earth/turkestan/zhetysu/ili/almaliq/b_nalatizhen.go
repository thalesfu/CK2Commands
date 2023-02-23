package almaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 那拉提镇NalatizhenBarony struct {
	feud.BaseBarony
}

var BNalatizhen那拉提镇 feud.Barony = &那拉提镇NalatizhenBarony{}

func init() {
	f := BNalatizhen那拉提镇.(*那拉提镇NalatizhenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nalatizhen",
		TitleName: "那拉提镇",
		TitleCode: "b_nalatizhen",
	}
}
