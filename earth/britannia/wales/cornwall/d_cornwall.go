package cornwall

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/cornwall/cornwall"
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/cornwall/devon"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CornwallDuke interface {
	feud.Duke
	CCornwall康沃尔() cornwall.CornwallCounty
	CDevon德文() devon.DevonCounty
}

type 康沃尔CornwallDuke struct {
	feud.BaseDuke
	康沃尔Cornwall cornwall.CornwallCounty
	德文Devon     devon.DevonCounty
}

func (d *康沃尔CornwallDuke) CCornwall康沃尔() cornwall.CornwallCounty {
	return d.康沃尔Cornwall
}

func (d *康沃尔CornwallDuke) CDevon德文() devon.DevonCounty {
	return d.德文Devon
}

var DCornwall康沃尔 CornwallDuke = &康沃尔CornwallDuke{}

func init() {
	f := DCornwall康沃尔.(*康沃尔CornwallDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "cornwall",
		TitleName: "康沃尔",
		TitleCode: "d_cornwall",
		Counties:  map[string]feud.County{},
	}

	f.康沃尔Cornwall = cornwall.CCornwall康沃尔
	f.康沃尔Cornwall.SetParent(f)

	f.德文Devon = devon.CDevon德文
	f.德文Devon.SetParent(f)

}
