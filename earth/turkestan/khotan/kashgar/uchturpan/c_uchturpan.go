package uchturpan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UchturpanCounty interface {
	feud.County
	BGaochetian高车田() feud.Barony
	BJizhuoguan济浊馆() feud.Barony
	BTunglingching东凌井() feud.Barony
	BUqturpan乌什() feud.Barony
	BYezheguan谒者馆() feud.Barony
	BYuyu郁于() feud.Barony
	BZhaojiazui赵家嘴() feud.Barony
}

type 乌什吐鲁番UchturpanCounty struct {
	feud.BaseCounty
	高车田Gaochetian    feud.Barony
	济浊馆Jizhuoguan    feud.Barony
	东凌井Tunglingching feud.Barony
	乌什Uqturpan       feud.Barony
	谒者馆Yezheguan     feud.Barony
	郁于Yuyu           feud.Barony
	赵家嘴Zhaojiazui    feud.Barony
}

func (c *乌什吐鲁番UchturpanCounty) BGaochetian高车田() feud.Barony {
	return c.高车田Gaochetian
}

func (c *乌什吐鲁番UchturpanCounty) BJizhuoguan济浊馆() feud.Barony {
	return c.济浊馆Jizhuoguan
}

func (c *乌什吐鲁番UchturpanCounty) BTunglingching东凌井() feud.Barony {
	return c.东凌井Tunglingching
}

func (c *乌什吐鲁番UchturpanCounty) BUqturpan乌什() feud.Barony {
	return c.乌什Uqturpan
}

func (c *乌什吐鲁番UchturpanCounty) BYezheguan谒者馆() feud.Barony {
	return c.谒者馆Yezheguan
}

func (c *乌什吐鲁番UchturpanCounty) BYuyu郁于() feud.Barony {
	return c.郁于Yuyu
}

func (c *乌什吐鲁番UchturpanCounty) BZhaojiazui赵家嘴() feud.Barony {
	return c.赵家嘴Zhaojiazui
}

var CUchturpan乌什吐鲁番 UchturpanCounty = &乌什吐鲁番UchturpanCounty{}

func init() {
	f := CUchturpan乌什吐鲁番.(*乌什吐鲁番UchturpanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1527",
		Title:     "uchturpan",
		TitleName: "乌什吐鲁番",
		TitleCode: "c_uchturpan",
		Baronies:  map[string]feud.Barony{},
	}

	f.高车田Gaochetian = BGaochetian高车田
	f.高车田Gaochetian.SetParent(f)

	f.济浊馆Jizhuoguan = BJizhuoguan济浊馆
	f.济浊馆Jizhuoguan.SetParent(f)

	f.东凌井Tunglingching = BTunglingching东凌井
	f.东凌井Tunglingching.SetParent(f)

	f.乌什Uqturpan = BUqturpan乌什
	f.乌什Uqturpan.SetParent(f)

	f.谒者馆Yezheguan = BYezheguan谒者馆
	f.谒者馆Yezheguan.SetParent(f)

	f.郁于Yuyu = BYuyu郁于
	f.郁于Yuyu.SetParent(f)

	f.赵家嘴Zhaojiazui = BZhaojiazui赵家嘴
	f.赵家嘴Zhaojiazui.SetParent(f)

}
