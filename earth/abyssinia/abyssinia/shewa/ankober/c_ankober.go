package ankober

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AnkoberCounty interface {
    feud.County
    BAliyuamba阿利尤安巴() 	feud.Barony
    BAnkober安科贝尔() 	feud.Barony
    BBerehet贝尔赫特() 	feud.Barony
    BDoqaqit多喀奇() 	feud.Barony
    BKeyagebriel克亚吉比利() 	feud.Barony
    BQundi赞迪() 	feud.Barony
    BSaqqa萨卡拉() 	feud.Barony
    BTegulet特古勒() 	feud.Barony
}

type 安科贝尔AnkoberCounty struct {
	feud.BaseCounty
	阿利尤安巴Aliyuamba 	feud.Barony
	安科贝尔Ankober 	feud.Barony
	贝尔赫特Berehet 	feud.Barony
	多喀奇Doqaqit 	feud.Barony
	克亚吉比利Keyagebriel 	feud.Barony
	赞迪Qundi 	feud.Barony
	萨卡拉Saqqa 	feud.Barony
	特古勒Tegulet 	feud.Barony
}

func (c *安科贝尔AnkoberCounty) BAliyuamba阿利尤安巴() feud.Barony {
	return c.阿利尤安巴Aliyuamba
}
    
func (c *安科贝尔AnkoberCounty) BAnkober安科贝尔() feud.Barony {
	return c.安科贝尔Ankober
}
    
func (c *安科贝尔AnkoberCounty) BBerehet贝尔赫特() feud.Barony {
	return c.贝尔赫特Berehet
}
    
func (c *安科贝尔AnkoberCounty) BDoqaqit多喀奇() feud.Barony {
	return c.多喀奇Doqaqit
}
    
func (c *安科贝尔AnkoberCounty) BKeyagebriel克亚吉比利() feud.Barony {
	return c.克亚吉比利Keyagebriel
}
    
func (c *安科贝尔AnkoberCounty) BQundi赞迪() feud.Barony {
	return c.赞迪Qundi
}
    
func (c *安科贝尔AnkoberCounty) BSaqqa萨卡拉() feud.Barony {
	return c.萨卡拉Saqqa
}
    
func (c *安科贝尔AnkoberCounty) BTegulet特古勒() feud.Barony {
	return c.特古勒Tegulet
}
    
var CAnkober安科贝尔 AnkoberCounty = &安科贝尔AnkoberCounty{}

func init() {
	f := CAnkober安科贝尔.(*安科贝尔AnkoberCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "882",
		Title:     "ankober",
		TitleName: "安科贝尔",
		TitleCode: "c_ankober",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿利尤安巴Aliyuamba = BAliyuamba阿利尤安巴
	f.阿利尤安巴Aliyuamba.SetParent(f)
	
	f.安科贝尔Ankober = BAnkober安科贝尔
	f.安科贝尔Ankober.SetParent(f)
	
	f.贝尔赫特Berehet = BBerehet贝尔赫特
	f.贝尔赫特Berehet.SetParent(f)
	
	f.多喀奇Doqaqit = BDoqaqit多喀奇
	f.多喀奇Doqaqit.SetParent(f)
	
	f.克亚吉比利Keyagebriel = BKeyagebriel克亚吉比利
	f.克亚吉比利Keyagebriel.SetParent(f)
	
	f.赞迪Qundi = BQundi赞迪
	f.赞迪Qundi.SetParent(f)
	
	f.萨卡拉Saqqa = BSaqqa萨卡拉
	f.萨卡拉Saqqa.SetParent(f)
	
	f.特古勒Tegulet = BTegulet特古勒
	f.特古勒Tegulet.SetParent(f)
	
}
