package penugonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PenugondaCounty interface {
    feud.County
    BAnantapur阿难多补罗() 	feud.Barony
    BGurramkonda具罗摩军荼() 	feud.Barony
    BLepakshi莱帕科什() 	feud.Barony
    BPenugonda卑奴郡荼() 	feud.Barony
    BRayadurgam拉亚德加姆() 	feud.Barony
    BSuragiri首罗耆厘() 	feud.Barony
    BTogarakunta都迦罗军吒() 	feud.Barony
    BUravakonda优罗婆军荼() 	feud.Barony
}

type 卑奴郡荼PenugondaCounty struct {
	feud.BaseCounty
	阿难多补罗Anantapur 	feud.Barony
	具罗摩军荼Gurramkonda 	feud.Barony
	莱帕科什Lepakshi 	feud.Barony
	卑奴郡荼Penugonda 	feud.Barony
	拉亚德加姆Rayadurgam 	feud.Barony
	首罗耆厘Suragiri 	feud.Barony
	都迦罗军吒Togarakunta 	feud.Barony
	优罗婆军荼Uravakonda 	feud.Barony
}

func (c *卑奴郡荼PenugondaCounty) BAnantapur阿难多补罗() feud.Barony {
	return c.阿难多补罗Anantapur
}
    
func (c *卑奴郡荼PenugondaCounty) BGurramkonda具罗摩军荼() feud.Barony {
	return c.具罗摩军荼Gurramkonda
}
    
func (c *卑奴郡荼PenugondaCounty) BLepakshi莱帕科什() feud.Barony {
	return c.莱帕科什Lepakshi
}
    
func (c *卑奴郡荼PenugondaCounty) BPenugonda卑奴郡荼() feud.Barony {
	return c.卑奴郡荼Penugonda
}
    
func (c *卑奴郡荼PenugondaCounty) BRayadurgam拉亚德加姆() feud.Barony {
	return c.拉亚德加姆Rayadurgam
}
    
func (c *卑奴郡荼PenugondaCounty) BSuragiri首罗耆厘() feud.Barony {
	return c.首罗耆厘Suragiri
}
    
func (c *卑奴郡荼PenugondaCounty) BTogarakunta都迦罗军吒() feud.Barony {
	return c.都迦罗军吒Togarakunta
}
    
func (c *卑奴郡荼PenugondaCounty) BUravakonda优罗婆军荼() feud.Barony {
	return c.优罗婆军荼Uravakonda
}
    
var CPenugonda卑奴郡荼 PenugondaCounty = &卑奴郡荼PenugondaCounty{}

func init() {
	f := CPenugonda卑奴郡荼.(*卑奴郡荼PenugondaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1199",
		Title:     "penugonda",
		TitleName: "卑奴郡荼",
		TitleCode: "c_penugonda",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿难多补罗Anantapur = BAnantapur阿难多补罗
	f.阿难多补罗Anantapur.SetParent(f)
	
	f.具罗摩军荼Gurramkonda = BGurramkonda具罗摩军荼
	f.具罗摩军荼Gurramkonda.SetParent(f)
	
	f.莱帕科什Lepakshi = BLepakshi莱帕科什
	f.莱帕科什Lepakshi.SetParent(f)
	
	f.卑奴郡荼Penugonda = BPenugonda卑奴郡荼
	f.卑奴郡荼Penugonda.SetParent(f)
	
	f.拉亚德加姆Rayadurgam = BRayadurgam拉亚德加姆
	f.拉亚德加姆Rayadurgam.SetParent(f)
	
	f.首罗耆厘Suragiri = BSuragiri首罗耆厘
	f.首罗耆厘Suragiri.SetParent(f)
	
	f.都迦罗军吒Togarakunta = BTogarakunta都迦罗军吒
	f.都迦罗军吒Togarakunta.SetParent(f)
	
	f.优罗婆军荼Uravakonda = BUravakonda优罗婆军荼
	f.优罗婆军荼Uravakonda.SetParent(f)
	
}
