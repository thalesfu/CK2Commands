package orania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OraniaCounty interface {
	feud.County
	BAinelberd艾因拜尔德() feud.Barony
	BAintemouchent艾因蒂姆沈特() feud.Barony
	BBenisaf贝尼萨夫() feud.Barony
	BGdyel格迪耶勒() feud.Barony
	BMaghnia马格尼耶() feud.Barony
	BMerselhadjad哈贾德() feud.Barony
	BMostaganem莫斯塔加纳姆() feud.Barony
	BOran奥兰() feud.Barony
}

type 奥兰OraniaCounty struct {
	feud.BaseCounty
	艾因拜尔德Ainelberd      feud.Barony
	艾因蒂姆沈特Aintemouchent feud.Barony
	贝尼萨夫Benisaf         feud.Barony
	格迪耶勒Gdyel           feud.Barony
	马格尼耶Maghnia         feud.Barony
	哈贾德Merselhadjad     feud.Barony
	莫斯塔加纳姆Mostaganem    feud.Barony
	奥兰Oran              feud.Barony
}

func (c *奥兰OraniaCounty) BAinelberd艾因拜尔德() feud.Barony {
	return c.艾因拜尔德Ainelberd
}

func (c *奥兰OraniaCounty) BAintemouchent艾因蒂姆沈特() feud.Barony {
	return c.艾因蒂姆沈特Aintemouchent
}

func (c *奥兰OraniaCounty) BBenisaf贝尼萨夫() feud.Barony {
	return c.贝尼萨夫Benisaf
}

func (c *奥兰OraniaCounty) BGdyel格迪耶勒() feud.Barony {
	return c.格迪耶勒Gdyel
}

func (c *奥兰OraniaCounty) BMaghnia马格尼耶() feud.Barony {
	return c.马格尼耶Maghnia
}

func (c *奥兰OraniaCounty) BMerselhadjad哈贾德() feud.Barony {
	return c.哈贾德Merselhadjad
}

func (c *奥兰OraniaCounty) BMostaganem莫斯塔加纳姆() feud.Barony {
	return c.莫斯塔加纳姆Mostaganem
}

func (c *奥兰OraniaCounty) BOran奥兰() feud.Barony {
	return c.奥兰Oran
}

var COrania奥兰 OraniaCounty = &奥兰OraniaCounty{}

func init() {
	f := COrania奥兰.(*奥兰OraniaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "832",
		Title:     "orania",
		TitleName: "奥兰",
		TitleCode: "c_orania",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾因拜尔德Ainelberd = BAinelberd艾因拜尔德
	f.艾因拜尔德Ainelberd.SetParent(f)

	f.艾因蒂姆沈特Aintemouchent = BAintemouchent艾因蒂姆沈特
	f.艾因蒂姆沈特Aintemouchent.SetParent(f)

	f.贝尼萨夫Benisaf = BBenisaf贝尼萨夫
	f.贝尼萨夫Benisaf.SetParent(f)

	f.格迪耶勒Gdyel = BGdyel格迪耶勒
	f.格迪耶勒Gdyel.SetParent(f)

	f.马格尼耶Maghnia = BMaghnia马格尼耶
	f.马格尼耶Maghnia.SetParent(f)

	f.哈贾德Merselhadjad = BMerselhadjad哈贾德
	f.哈贾德Merselhadjad.SetParent(f)

	f.莫斯塔加纳姆Mostaganem = BMostaganem莫斯塔加纳姆
	f.莫斯塔加纳姆Mostaganem.SetParent(f)

	f.奥兰Oran = BOran奥兰
	f.奥兰Oran.SetParent(f)

}
