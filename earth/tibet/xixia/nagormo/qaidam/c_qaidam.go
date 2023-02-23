package qaidam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QaidamCounty interface {
	feud.County
	BAolaohe嗷唠河() feud.Barony
	BDaqaidam大柴旦() feud.Barony
	BLigouchu里沟渠() feud.Barony
	BQaidam柴旦() feud.Barony
	BXiaochaidamu小柴达木() feud.Barony
	BXitishan锡铁山() feud.Barony
	BYinmaxia饮马峡() feud.Barony
}

type 柴旦QaidamCounty struct {
	feud.BaseCounty
	嗷唠河Aolaohe       feud.Barony
	大柴旦Daqaidam      feud.Barony
	里沟渠Ligouchu      feud.Barony
	柴旦Qaidam         feud.Barony
	小柴达木Xiaochaidamu feud.Barony
	锡铁山Xitishan      feud.Barony
	饮马峡Yinmaxia      feud.Barony
}

func (c *柴旦QaidamCounty) BAolaohe嗷唠河() feud.Barony {
	return c.嗷唠河Aolaohe
}

func (c *柴旦QaidamCounty) BDaqaidam大柴旦() feud.Barony {
	return c.大柴旦Daqaidam
}

func (c *柴旦QaidamCounty) BLigouchu里沟渠() feud.Barony {
	return c.里沟渠Ligouchu
}

func (c *柴旦QaidamCounty) BQaidam柴旦() feud.Barony {
	return c.柴旦Qaidam
}

func (c *柴旦QaidamCounty) BXiaochaidamu小柴达木() feud.Barony {
	return c.小柴达木Xiaochaidamu
}

func (c *柴旦QaidamCounty) BXitishan锡铁山() feud.Barony {
	return c.锡铁山Xitishan
}

func (c *柴旦QaidamCounty) BYinmaxia饮马峡() feud.Barony {
	return c.饮马峡Yinmaxia
}

var CQaidam柴旦 QaidamCounty = &柴旦QaidamCounty{}

func init() {
	f := CQaidam柴旦.(*柴旦QaidamCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1570",
		Title:     "qaidam",
		TitleName: "柴旦",
		TitleCode: "c_qaidam",
		Baronies:  map[string]feud.Barony{},
	}

	f.嗷唠河Aolaohe = BAolaohe嗷唠河
	f.嗷唠河Aolaohe.SetParent(f)

	f.大柴旦Daqaidam = BDaqaidam大柴旦
	f.大柴旦Daqaidam.SetParent(f)

	f.里沟渠Ligouchu = BLigouchu里沟渠
	f.里沟渠Ligouchu.SetParent(f)

	f.柴旦Qaidam = BQaidam柴旦
	f.柴旦Qaidam.SetParent(f)

	f.小柴达木Xiaochaidamu = BXiaochaidamu小柴达木
	f.小柴达木Xiaochaidamu.SetParent(f)

	f.锡铁山Xitishan = BXitishan锡铁山
	f.锡铁山Xitishan.SetParent(f)

	f.饮马峡Yinmaxia = BYinmaxia饮马峡
	f.饮马峡Yinmaxia.SetParent(f)

}
