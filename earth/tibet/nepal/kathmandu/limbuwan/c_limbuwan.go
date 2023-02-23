package limbuwan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LimbuwanCounty interface {
	feud.County
	BDhankuta驮那矩吒() feud.Barony
	BGograha瞿揭罗诃() feud.Barony
	BMorong牧龙() feud.Barony
	BMyanglung迈阿格朗德() feud.Barony
	BTambar耽婆罗() feud.Barony
	BTaplejung达布莱宗() feud.Barony
}

type 林布旺LimbuwanCounty struct {
	feud.BaseCounty
	驮那矩吒Dhankuta   feud.Barony
	瞿揭罗诃Gograha    feud.Barony
	牧龙Morong       feud.Barony
	迈阿格朗德Myanglung feud.Barony
	耽婆罗Tambar      feud.Barony
	达布莱宗Taplejung  feud.Barony
}

func (c *林布旺LimbuwanCounty) BDhankuta驮那矩吒() feud.Barony {
	return c.驮那矩吒Dhankuta
}

func (c *林布旺LimbuwanCounty) BGograha瞿揭罗诃() feud.Barony {
	return c.瞿揭罗诃Gograha
}

func (c *林布旺LimbuwanCounty) BMorong牧龙() feud.Barony {
	return c.牧龙Morong
}

func (c *林布旺LimbuwanCounty) BMyanglung迈阿格朗德() feud.Barony {
	return c.迈阿格朗德Myanglung
}

func (c *林布旺LimbuwanCounty) BTambar耽婆罗() feud.Barony {
	return c.耽婆罗Tambar
}

func (c *林布旺LimbuwanCounty) BTaplejung达布莱宗() feud.Barony {
	return c.达布莱宗Taplejung
}

var CLimbuwan林布旺 LimbuwanCounty = &林布旺LimbuwanCounty{}

func init() {
	f := CLimbuwan林布旺.(*林布旺LimbuwanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1479",
		Title:     "limbuwan",
		TitleName: "林布旺",
		TitleCode: "c_limbuwan",
		Baronies:  map[string]feud.Barony{},
	}

	f.驮那矩吒Dhankuta = BDhankuta驮那矩吒
	f.驮那矩吒Dhankuta.SetParent(f)

	f.瞿揭罗诃Gograha = BGograha瞿揭罗诃
	f.瞿揭罗诃Gograha.SetParent(f)

	f.牧龙Morong = BMorong牧龙
	f.牧龙Morong.SetParent(f)

	f.迈阿格朗德Myanglung = BMyanglung迈阿格朗德
	f.迈阿格朗德Myanglung.SetParent(f)

	f.耽婆罗Tambar = BTambar耽婆罗
	f.耽婆罗Tambar.SetParent(f)

	f.达布莱宗Taplejung = BTaplejung达布莱宗
	f.达布莱宗Taplejung.SetParent(f)

}
