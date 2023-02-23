package beersheb

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BeershebCounty interface {
	feud.County
	BAbuzqayqa阿布扎奇亚拉() feud.Barony
	BBeersheb贝尔谢巴() feud.Barony
	BDebir底璧() feud.Barony
	BEstemon埃斯特莫() feud.Barony
	BGilat吉来特() feud.Barony
	BMassada马萨达() feud.Barony
	BOfakim奥法基姆() feud.Barony
	BRahat拉哈特() feud.Barony
}

type 贝尔谢巴BeershebCounty struct {
	feud.BaseCounty
	阿布扎奇亚拉Abuzqayqa feud.Barony
	贝尔谢巴Beersheb    feud.Barony
	底璧Debir         feud.Barony
	埃斯特莫Estemon     feud.Barony
	吉来特Gilat        feud.Barony
	马萨达Massada      feud.Barony
	奥法基姆Ofakim      feud.Barony
	拉哈特Rahat        feud.Barony
}

func (c *贝尔谢巴BeershebCounty) BAbuzqayqa阿布扎奇亚拉() feud.Barony {
	return c.阿布扎奇亚拉Abuzqayqa
}

func (c *贝尔谢巴BeershebCounty) BBeersheb贝尔谢巴() feud.Barony {
	return c.贝尔谢巴Beersheb
}

func (c *贝尔谢巴BeershebCounty) BDebir底璧() feud.Barony {
	return c.底璧Debir
}

func (c *贝尔谢巴BeershebCounty) BEstemon埃斯特莫() feud.Barony {
	return c.埃斯特莫Estemon
}

func (c *贝尔谢巴BeershebCounty) BGilat吉来特() feud.Barony {
	return c.吉来特Gilat
}

func (c *贝尔谢巴BeershebCounty) BMassada马萨达() feud.Barony {
	return c.马萨达Massada
}

func (c *贝尔谢巴BeershebCounty) BOfakim奥法基姆() feud.Barony {
	return c.奥法基姆Ofakim
}

func (c *贝尔谢巴BeershebCounty) BRahat拉哈特() feud.Barony {
	return c.拉哈特Rahat
}

var CBeersheb贝尔谢巴 BeershebCounty = &贝尔谢巴BeershebCounty{}

func init() {
	f := CBeersheb贝尔谢巴.(*贝尔谢巴BeershebCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "779",
		Title:     "beersheb",
		TitleName: "贝尔谢巴",
		TitleCode: "c_beersheb",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿布扎奇亚拉Abuzqayqa = BAbuzqayqa阿布扎奇亚拉
	f.阿布扎奇亚拉Abuzqayqa.SetParent(f)

	f.贝尔谢巴Beersheb = BBeersheb贝尔谢巴
	f.贝尔谢巴Beersheb.SetParent(f)

	f.底璧Debir = BDebir底璧
	f.底璧Debir.SetParent(f)

	f.埃斯特莫Estemon = BEstemon埃斯特莫
	f.埃斯特莫Estemon.SetParent(f)

	f.吉来特Gilat = BGilat吉来特
	f.吉来特Gilat.SetParent(f)

	f.马萨达Massada = BMassada马萨达
	f.马萨达Massada.SetParent(f)

	f.奥法基姆Ofakim = BOfakim奥法基姆
	f.奥法基姆Ofakim.SetParent(f)

	f.拉哈特Rahat = BRahat拉哈特
	f.拉哈特Rahat.SetParent(f)

}
