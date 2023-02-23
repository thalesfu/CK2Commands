package gyesar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GyesarCounty interface {
	feud.County
	BGyesar杰萨() feud.Barony
	BJyangling蒋陵() feud.Barony
	BKarongshar嘎荣下() feud.Barony
	BKhongtog贡朵() feud.Barony
	BPashul帕苏尔() feud.Barony
	BSemola桑木拉() feud.Barony
	BTrengsham惩香() feud.Barony
}

type 杰萨GyesarCounty struct {
	feud.BaseCounty
	杰萨Gyesar      feud.Barony
	蒋陵Jyangling   feud.Barony
	嘎荣下Karongshar feud.Barony
	贡朵Khongtog    feud.Barony
	帕苏尔Pashul     feud.Barony
	桑木拉Semola     feud.Barony
	惩香Trengsham   feud.Barony
}

func (c *杰萨GyesarCounty) BGyesar杰萨() feud.Barony {
	return c.杰萨Gyesar
}

func (c *杰萨GyesarCounty) BJyangling蒋陵() feud.Barony {
	return c.蒋陵Jyangling
}

func (c *杰萨GyesarCounty) BKarongshar嘎荣下() feud.Barony {
	return c.嘎荣下Karongshar
}

func (c *杰萨GyesarCounty) BKhongtog贡朵() feud.Barony {
	return c.贡朵Khongtog
}

func (c *杰萨GyesarCounty) BPashul帕苏尔() feud.Barony {
	return c.帕苏尔Pashul
}

func (c *杰萨GyesarCounty) BSemola桑木拉() feud.Barony {
	return c.桑木拉Semola
}

func (c *杰萨GyesarCounty) BTrengsham惩香() feud.Barony {
	return c.惩香Trengsham
}

var CGyesar杰萨 GyesarCounty = &杰萨GyesarCounty{}

func init() {
	f := CGyesar杰萨.(*杰萨GyesarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1566",
		Title:     "gyesar",
		TitleName: "杰萨",
		TitleCode: "c_gyesar",
		Baronies:  map[string]feud.Barony{},
	}

	f.杰萨Gyesar = BGyesar杰萨
	f.杰萨Gyesar.SetParent(f)

	f.蒋陵Jyangling = BJyangling蒋陵
	f.蒋陵Jyangling.SetParent(f)

	f.嘎荣下Karongshar = BKarongshar嘎荣下
	f.嘎荣下Karongshar.SetParent(f)

	f.贡朵Khongtog = BKhongtog贡朵
	f.贡朵Khongtog.SetParent(f)

	f.帕苏尔Pashul = BPashul帕苏尔
	f.帕苏尔Pashul.SetParent(f)

	f.桑木拉Semola = BSemola桑木拉
	f.桑木拉Semola.SetParent(f)

	f.惩香Trengsham = BTrengsham惩香
	f.惩香Trengsham.SetParent(f)

}
