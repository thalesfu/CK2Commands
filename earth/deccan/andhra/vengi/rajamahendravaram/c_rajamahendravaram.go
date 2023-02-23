package rajamahendravaram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RajamahendravaramCounty interface {
	feud.County
	BGetalsud祇呾罗窣陀() feud.Barony
	BHemagiri醯摩耆厘() feud.Barony
	BKirandul基兰杜尔() feud.Barony
	BKuwakonda鸠婆军荼() feud.Barony
	BPathyar波梯耶罗() feud.Barony
	BPattiseema帕迪瑟玛() feud.Barony
	BRajamahendravaram罗阇摩醯因陀罗伐蓝() feud.Barony
}

type 罗阇摩醯因陀罗伐蓝RajamahendravaramCounty struct {
	feud.BaseCounty
	祇呾罗窣陀Getalsud              feud.Barony
	醯摩耆厘Hemagiri               feud.Barony
	基兰杜尔Kirandul               feud.Barony
	鸠婆军荼Kuwakonda              feud.Barony
	波梯耶罗Pathyar                feud.Barony
	帕迪瑟玛Pattiseema             feud.Barony
	罗阇摩醯因陀罗伐蓝Rajamahendravaram feud.Barony
}

func (c *罗阇摩醯因陀罗伐蓝RajamahendravaramCounty) BGetalsud祇呾罗窣陀() feud.Barony {
	return c.祇呾罗窣陀Getalsud
}

func (c *罗阇摩醯因陀罗伐蓝RajamahendravaramCounty) BHemagiri醯摩耆厘() feud.Barony {
	return c.醯摩耆厘Hemagiri
}

func (c *罗阇摩醯因陀罗伐蓝RajamahendravaramCounty) BKirandul基兰杜尔() feud.Barony {
	return c.基兰杜尔Kirandul
}

func (c *罗阇摩醯因陀罗伐蓝RajamahendravaramCounty) BKuwakonda鸠婆军荼() feud.Barony {
	return c.鸠婆军荼Kuwakonda
}

func (c *罗阇摩醯因陀罗伐蓝RajamahendravaramCounty) BPathyar波梯耶罗() feud.Barony {
	return c.波梯耶罗Pathyar
}

func (c *罗阇摩醯因陀罗伐蓝RajamahendravaramCounty) BPattiseema帕迪瑟玛() feud.Barony {
	return c.帕迪瑟玛Pattiseema
}

func (c *罗阇摩醯因陀罗伐蓝RajamahendravaramCounty) BRajamahendravaram罗阇摩醯因陀罗伐蓝() feud.Barony {
	return c.罗阇摩醯因陀罗伐蓝Rajamahendravaram
}

var CRajamahendravaram罗阇摩醯因陀罗伐蓝 RajamahendravaramCounty = &罗阇摩醯因陀罗伐蓝RajamahendravaramCounty{}

func init() {
	f := CRajamahendravaram罗阇摩醯因陀罗伐蓝.(*罗阇摩醯因陀罗伐蓝RajamahendravaramCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1222",
		Title:     "rajamahendravaram",
		TitleName: "罗阇摩醯因陀罗伐蓝",
		TitleCode: "c_rajamahendravaram",
		Baronies:  map[string]feud.Barony{},
	}

	f.祇呾罗窣陀Getalsud = BGetalsud祇呾罗窣陀
	f.祇呾罗窣陀Getalsud.SetParent(f)

	f.醯摩耆厘Hemagiri = BHemagiri醯摩耆厘
	f.醯摩耆厘Hemagiri.SetParent(f)

	f.基兰杜尔Kirandul = BKirandul基兰杜尔
	f.基兰杜尔Kirandul.SetParent(f)

	f.鸠婆军荼Kuwakonda = BKuwakonda鸠婆军荼
	f.鸠婆军荼Kuwakonda.SetParent(f)

	f.波梯耶罗Pathyar = BPathyar波梯耶罗
	f.波梯耶罗Pathyar.SetParent(f)

	f.帕迪瑟玛Pattiseema = BPattiseema帕迪瑟玛
	f.帕迪瑟玛Pattiseema.SetParent(f)

	f.罗阇摩醯因陀罗伐蓝Rajamahendravaram = BRajamahendravaram罗阇摩醯因陀罗伐蓝
	f.罗阇摩醯因陀罗伐蓝Rajamahendravaram.SetParent(f)

}
