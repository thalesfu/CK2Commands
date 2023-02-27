package karvuna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡利斯特里亚SilistriaBarony struct {
	feud.BaseBarony
}

var BSilistria锡利斯特里亚 feud.Barony = &锡利斯特里亚SilistriaBarony{}

func init() {
    f := BSilistria锡利斯特里亚.(*锡利斯特里亚SilistriaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "silistria",
		TitleName: "锡利斯特里亚",
		TitleCode: "b_silistria",
	}
}
