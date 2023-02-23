package cuenca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CuencaCounty interface {
	feud.County
	BCuenca昆卡() feud.Barony
	BLaspedroneras拉斯佩德罗涅拉斯() feud.Barony
	BMotadelcuervo莫塔德尔奎尔沃() feud.Barony
	BSisante西桑特() feud.Barony
	BTarancon塔兰孔() feud.Barony
	BUcles乌克莱斯() feud.Barony
	BVillanuevadelajara哈拉新镇() feud.Barony
}

type 昆卡CuencaCounty struct {
	feud.BaseCounty
	昆卡Cuenca               feud.Barony
	拉斯佩德罗涅拉斯Laspedroneras  feud.Barony
	莫塔德尔奎尔沃Motadelcuervo   feud.Barony
	西桑特Sisante             feud.Barony
	塔兰孔Tarancon            feud.Barony
	乌克莱斯Ucles              feud.Barony
	哈拉新镇Villanuevadelajara feud.Barony
}

func (c *昆卡CuencaCounty) BCuenca昆卡() feud.Barony {
	return c.昆卡Cuenca
}

func (c *昆卡CuencaCounty) BLaspedroneras拉斯佩德罗涅拉斯() feud.Barony {
	return c.拉斯佩德罗涅拉斯Laspedroneras
}

func (c *昆卡CuencaCounty) BMotadelcuervo莫塔德尔奎尔沃() feud.Barony {
	return c.莫塔德尔奎尔沃Motadelcuervo
}

func (c *昆卡CuencaCounty) BSisante西桑特() feud.Barony {
	return c.西桑特Sisante
}

func (c *昆卡CuencaCounty) BTarancon塔兰孔() feud.Barony {
	return c.塔兰孔Tarancon
}

func (c *昆卡CuencaCounty) BUcles乌克莱斯() feud.Barony {
	return c.乌克莱斯Ucles
}

func (c *昆卡CuencaCounty) BVillanuevadelajara哈拉新镇() feud.Barony {
	return c.哈拉新镇Villanuevadelajara
}

var CCuenca昆卡 CuencaCounty = &昆卡CuencaCounty{}

func init() {
	f := CCuenca昆卡.(*昆卡CuencaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "177",
		Title:     "cuenca",
		TitleName: "昆卡",
		TitleCode: "c_cuenca",
		Baronies:  map[string]feud.Barony{},
	}

	f.昆卡Cuenca = BCuenca昆卡
	f.昆卡Cuenca.SetParent(f)

	f.拉斯佩德罗涅拉斯Laspedroneras = BLaspedroneras拉斯佩德罗涅拉斯
	f.拉斯佩德罗涅拉斯Laspedroneras.SetParent(f)

	f.莫塔德尔奎尔沃Motadelcuervo = BMotadelcuervo莫塔德尔奎尔沃
	f.莫塔德尔奎尔沃Motadelcuervo.SetParent(f)

	f.西桑特Sisante = BSisante西桑特
	f.西桑特Sisante.SetParent(f)

	f.塔兰孔Tarancon = BTarancon塔兰孔
	f.塔兰孔Tarancon.SetParent(f)

	f.乌克莱斯Ucles = BUcles乌克莱斯
	f.乌克莱斯Ucles.SetParent(f)

	f.哈拉新镇Villanuevadelajara = BVillanuevadelajara哈拉新镇
	f.哈拉新镇Villanuevadelajara.SetParent(f)

}
