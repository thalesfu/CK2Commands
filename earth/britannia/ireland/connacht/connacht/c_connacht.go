package connacht

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ConnachtCounty interface {
	feud.County
	BAnchory阿克雷() feud.Barony
	BClonfert康菲特() feud.Barony
	BElphin埃尔芬() feud.Barony
	BGalway戈尔韦() feud.Barony
	BKillala基拉拉() feud.Barony
	BRoscommon罗斯康芒() feud.Barony
	BTuam蒂厄姆() feud.Barony
}

type 西康诺特ConnachtCounty struct {
	feud.BaseCounty
	阿克雷Anchory    feud.Barony
	康菲特Clonfert   feud.Barony
	埃尔芬Elphin     feud.Barony
	戈尔韦Galway     feud.Barony
	基拉拉Killala    feud.Barony
	罗斯康芒Roscommon feud.Barony
	蒂厄姆Tuam       feud.Barony
}

func (c *西康诺特ConnachtCounty) BAnchory阿克雷() feud.Barony {
	return c.阿克雷Anchory
}

func (c *西康诺特ConnachtCounty) BClonfert康菲特() feud.Barony {
	return c.康菲特Clonfert
}

func (c *西康诺特ConnachtCounty) BElphin埃尔芬() feud.Barony {
	return c.埃尔芬Elphin
}

func (c *西康诺特ConnachtCounty) BGalway戈尔韦() feud.Barony {
	return c.戈尔韦Galway
}

func (c *西康诺特ConnachtCounty) BKillala基拉拉() feud.Barony {
	return c.基拉拉Killala
}

func (c *西康诺特ConnachtCounty) BRoscommon罗斯康芒() feud.Barony {
	return c.罗斯康芒Roscommon
}

func (c *西康诺特ConnachtCounty) BTuam蒂厄姆() feud.Barony {
	return c.蒂厄姆Tuam
}

var CConnacht西康诺特 ConnachtCounty = &西康诺特ConnachtCounty{}

func init() {
	f := CConnacht西康诺特.(*西康诺特ConnachtCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "9",
		Title:     "connacht",
		TitleName: "西康诺特",
		TitleCode: "c_connacht",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克雷Anchory = BAnchory阿克雷
	f.阿克雷Anchory.SetParent(f)

	f.康菲特Clonfert = BClonfert康菲特
	f.康菲特Clonfert.SetParent(f)

	f.埃尔芬Elphin = BElphin埃尔芬
	f.埃尔芬Elphin.SetParent(f)

	f.戈尔韦Galway = BGalway戈尔韦
	f.戈尔韦Galway.SetParent(f)

	f.基拉拉Killala = BKillala基拉拉
	f.基拉拉Killala.SetParent(f)

	f.罗斯康芒Roscommon = BRoscommon罗斯康芒
	f.罗斯康芒Roscommon.SetParent(f)

	f.蒂厄姆Tuam = BTuam蒂厄姆
	f.蒂厄姆Tuam.SetParent(f)

}
