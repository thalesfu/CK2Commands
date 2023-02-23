package albania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AlbaniaCounty interface {
	feud.County
	BBarda巴尔达() feud.Barony
	BChukhurkabala楚霍尔卡巴拉() feud.Barony
	BDarussoltan达乌斯索坦() feud.Barony
	BEmli叶木里() feud.Barony
	BGanja占贾() feud.Barony
	BGelersengorersen格雷森格拉森() feud.Barony
	BKabala卡巴莱() feud.Barony
	BShaki沙阿基() feud.Barony
}

type 阿尔巴尼亚AlbaniaCounty struct {
	feud.BaseCounty
	巴尔达Barda               feud.Barony
	楚霍尔卡巴拉Chukhurkabala    feud.Barony
	达乌斯索坦Darussoltan       feud.Barony
	叶木里Emli                feud.Barony
	占贾Ganja                feud.Barony
	格雷森格拉森Gelersengorersen feud.Barony
	卡巴莱Kabala              feud.Barony
	沙阿基Shaki               feud.Barony
}

func (c *阿尔巴尼亚AlbaniaCounty) BBarda巴尔达() feud.Barony {
	return c.巴尔达Barda
}

func (c *阿尔巴尼亚AlbaniaCounty) BChukhurkabala楚霍尔卡巴拉() feud.Barony {
	return c.楚霍尔卡巴拉Chukhurkabala
}

func (c *阿尔巴尼亚AlbaniaCounty) BDarussoltan达乌斯索坦() feud.Barony {
	return c.达乌斯索坦Darussoltan
}

func (c *阿尔巴尼亚AlbaniaCounty) BEmli叶木里() feud.Barony {
	return c.叶木里Emli
}

func (c *阿尔巴尼亚AlbaniaCounty) BGanja占贾() feud.Barony {
	return c.占贾Ganja
}

func (c *阿尔巴尼亚AlbaniaCounty) BGelersengorersen格雷森格拉森() feud.Barony {
	return c.格雷森格拉森Gelersengorersen
}

func (c *阿尔巴尼亚AlbaniaCounty) BKabala卡巴莱() feud.Barony {
	return c.卡巴莱Kabala
}

func (c *阿尔巴尼亚AlbaniaCounty) BShaki沙阿基() feud.Barony {
	return c.沙阿基Shaki
}

var CAlbania阿尔巴尼亚 AlbaniaCounty = &阿尔巴尼亚AlbaniaCounty{}

func init() {
	f := CAlbania阿尔巴尼亚.(*阿尔巴尼亚AlbaniaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "673",
		Title:     "albania",
		TitleName: "阿尔巴尼亚",
		TitleCode: "c_albania",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尔达Barda = BBarda巴尔达
	f.巴尔达Barda.SetParent(f)

	f.楚霍尔卡巴拉Chukhurkabala = BChukhurkabala楚霍尔卡巴拉
	f.楚霍尔卡巴拉Chukhurkabala.SetParent(f)

	f.达乌斯索坦Darussoltan = BDarussoltan达乌斯索坦
	f.达乌斯索坦Darussoltan.SetParent(f)

	f.叶木里Emli = BEmli叶木里
	f.叶木里Emli.SetParent(f)

	f.占贾Ganja = BGanja占贾
	f.占贾Ganja.SetParent(f)

	f.格雷森格拉森Gelersengorersen = BGelersengorersen格雷森格拉森
	f.格雷森格拉森Gelersengorersen.SetParent(f)

	f.卡巴莱Kabala = BKabala卡巴莱
	f.卡巴莱Kabala.SetParent(f)

	f.沙阿基Shaki = BShaki沙阿基
	f.沙阿基Shaki.SetParent(f)

}
