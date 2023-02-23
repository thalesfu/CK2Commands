package finnmark

import (
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sapmi/finnmark/finnmark"
	"github.com/thalesfu/CK2Commands/earth/scandinavia/sapmi/finnmark/nordland"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FinnmarkDuke interface {
	feud.Duke
	CFinnmark芬马克() finnmark.FinnmarkCounty
	CNordland诺德兰() nordland.NordlandCounty
}

type 芬马克FinnmarkDuke struct {
	feud.BaseDuke
	芬马克Finnmark finnmark.FinnmarkCounty
	诺德兰Nordland nordland.NordlandCounty
}

func (d *芬马克FinnmarkDuke) CFinnmark芬马克() finnmark.FinnmarkCounty {
	return d.芬马克Finnmark
}

func (d *芬马克FinnmarkDuke) CNordland诺德兰() nordland.NordlandCounty {
	return d.诺德兰Nordland
}

var DFinnmark芬马克 FinnmarkDuke = &芬马克FinnmarkDuke{}

func init() {
	f := DFinnmark芬马克.(*芬马克FinnmarkDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "finnmark",
		TitleName: "芬马克",
		TitleCode: "d_finnmark",
		Counties:  map[string]feud.County{},
	}

	f.芬马克Finnmark = finnmark.CFinnmark芬马克
	f.芬马克Finnmark.SetParent(f)

	f.诺德兰Nordland = nordland.CNordland诺德兰
	f.诺德兰Nordland.SetParent(f)

}
