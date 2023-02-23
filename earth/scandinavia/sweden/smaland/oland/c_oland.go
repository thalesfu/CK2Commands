package oland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OlandCounty interface {
	feud.County
	BAlgutsrum奥古斯朗() feud.Barony
	BBorgholm博里霍尔姆() feud.Barony
	BGardby戈德比() feud.Barony
	BGlomminge格勒明厄() feud.Barony
	BGronhogen格伦赫根() feud.Barony
	BHulterstad胡尔特斯塔德() feud.Barony
	BKopingsvik雪平斯维克() feud.Barony
	BMykleby米克雷比() feud.Barony
	BOttenby乌滕比() feud.Barony
}

type 厄兰OlandCounty struct {
	feud.BaseCounty
	奥古斯朗Algutsrum    feud.Barony
	博里霍尔姆Borgholm    feud.Barony
	戈德比Gardby        feud.Barony
	格勒明厄Glomminge    feud.Barony
	格伦赫根Gronhogen    feud.Barony
	胡尔特斯塔德Hulterstad feud.Barony
	雪平斯维克Kopingsvik  feud.Barony
	米克雷比Mykleby      feud.Barony
	乌滕比Ottenby       feud.Barony
}

func (c *厄兰OlandCounty) BAlgutsrum奥古斯朗() feud.Barony {
	return c.奥古斯朗Algutsrum
}

func (c *厄兰OlandCounty) BBorgholm博里霍尔姆() feud.Barony {
	return c.博里霍尔姆Borgholm
}

func (c *厄兰OlandCounty) BGardby戈德比() feud.Barony {
	return c.戈德比Gardby
}

func (c *厄兰OlandCounty) BGlomminge格勒明厄() feud.Barony {
	return c.格勒明厄Glomminge
}

func (c *厄兰OlandCounty) BGronhogen格伦赫根() feud.Barony {
	return c.格伦赫根Gronhogen
}

func (c *厄兰OlandCounty) BHulterstad胡尔特斯塔德() feud.Barony {
	return c.胡尔特斯塔德Hulterstad
}

func (c *厄兰OlandCounty) BKopingsvik雪平斯维克() feud.Barony {
	return c.雪平斯维克Kopingsvik
}

func (c *厄兰OlandCounty) BMykleby米克雷比() feud.Barony {
	return c.米克雷比Mykleby
}

func (c *厄兰OlandCounty) BOttenby乌滕比() feud.Barony {
	return c.乌滕比Ottenby
}

var COland厄兰 OlandCounty = &厄兰OlandCounty{}

func init() {
	f := COland厄兰.(*厄兰OlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "300",
		Title:     "oland",
		TitleName: "厄兰",
		TitleCode: "c_oland",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥古斯朗Algutsrum = BAlgutsrum奥古斯朗
	f.奥古斯朗Algutsrum.SetParent(f)

	f.博里霍尔姆Borgholm = BBorgholm博里霍尔姆
	f.博里霍尔姆Borgholm.SetParent(f)

	f.戈德比Gardby = BGardby戈德比
	f.戈德比Gardby.SetParent(f)

	f.格勒明厄Glomminge = BGlomminge格勒明厄
	f.格勒明厄Glomminge.SetParent(f)

	f.格伦赫根Gronhogen = BGronhogen格伦赫根
	f.格伦赫根Gronhogen.SetParent(f)

	f.胡尔特斯塔德Hulterstad = BHulterstad胡尔特斯塔德
	f.胡尔特斯塔德Hulterstad.SetParent(f)

	f.雪平斯维克Kopingsvik = BKopingsvik雪平斯维克
	f.雪平斯维克Kopingsvik.SetParent(f)

	f.米克雷比Mykleby = BMykleby米克雷比
	f.米克雷比Mykleby.SetParent(f)

	f.乌滕比Ottenby = BOttenby乌滕比
	f.乌滕比Ottenby.SetParent(f)

}
