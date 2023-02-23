package ob

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ObCounty interface {
	feud.County
	BAkutikha阿库季哈() feud.Barony
	BBarnaul巴尔瑙尔() feud.Barony
	BBuranovo布拉诺沃() feud.Barony
	BChekanikha切卡尼哈() feud.Barony
	BOb鄂毕() feud.Barony
	BOdintsovka奥金佐夫卡() feud.Barony
	BVolodarka沃洛达尔卡() feud.Barony
}

type 鄂毕ObCounty struct {
	feud.BaseCounty
	阿库季哈Akutikha    feud.Barony
	巴尔瑙尔Barnaul     feud.Barony
	布拉诺沃Buranovo    feud.Barony
	切卡尼哈Chekanikha  feud.Barony
	鄂毕Ob            feud.Barony
	奥金佐夫卡Odintsovka feud.Barony
	沃洛达尔卡Volodarka  feud.Barony
}

func (c *鄂毕ObCounty) BAkutikha阿库季哈() feud.Barony {
	return c.阿库季哈Akutikha
}

func (c *鄂毕ObCounty) BBarnaul巴尔瑙尔() feud.Barony {
	return c.巴尔瑙尔Barnaul
}

func (c *鄂毕ObCounty) BBuranovo布拉诺沃() feud.Barony {
	return c.布拉诺沃Buranovo
}

func (c *鄂毕ObCounty) BChekanikha切卡尼哈() feud.Barony {
	return c.切卡尼哈Chekanikha
}

func (c *鄂毕ObCounty) BOb鄂毕() feud.Barony {
	return c.鄂毕Ob
}

func (c *鄂毕ObCounty) BOdintsovka奥金佐夫卡() feud.Barony {
	return c.奥金佐夫卡Odintsovka
}

func (c *鄂毕ObCounty) BVolodarka沃洛达尔卡() feud.Barony {
	return c.沃洛达尔卡Volodarka
}

var COb鄂毕 ObCounty = &鄂毕ObCounty{}

func init() {
	f := COb鄂毕.(*鄂毕ObCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1873",
		Title:     "ob",
		TitleName: "鄂毕",
		TitleCode: "c_ob",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿库季哈Akutikha = BAkutikha阿库季哈
	f.阿库季哈Akutikha.SetParent(f)

	f.巴尔瑙尔Barnaul = BBarnaul巴尔瑙尔
	f.巴尔瑙尔Barnaul.SetParent(f)

	f.布拉诺沃Buranovo = BBuranovo布拉诺沃
	f.布拉诺沃Buranovo.SetParent(f)

	f.切卡尼哈Chekanikha = BChekanikha切卡尼哈
	f.切卡尼哈Chekanikha.SetParent(f)

	f.鄂毕Ob = BOb鄂毕
	f.鄂毕Ob.SetParent(f)

	f.奥金佐夫卡Odintsovka = BOdintsovka奥金佐夫卡
	f.奥金佐夫卡Odintsovka.SetParent(f)

	f.沃洛达尔卡Volodarka = BVolodarka沃洛达尔卡
	f.沃洛达尔卡Volodarka.SetParent(f)

}
