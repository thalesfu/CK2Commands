package ashli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AshliCounty interface {
    feud.County
    BAshli阿什勒() 	feud.Barony
    BBlagovescen布拉戈维申() 	feud.Barony
    BBoro伯勒() 	feud.Barony
    BDaulakan代于莱坎() 	feud.Barony
    BDurtojle久尔秋利() 	feud.Barony
    BJanauyl亚纳乌尔() 	feud.Barony
    BTuimasy图伊马济() 	feud.Barony
    BYamantaw亚曼套() 	feud.Barony
}

type 阿什勒AshliCounty struct {
	feud.BaseCounty
	阿什勒Ashli 	feud.Barony
	布拉戈维申Blagovescen 	feud.Barony
	伯勒Boro 	feud.Barony
	代于莱坎Daulakan 	feud.Barony
	久尔秋利Durtojle 	feud.Barony
	亚纳乌尔Janauyl 	feud.Barony
	图伊马济Tuimasy 	feud.Barony
	亚曼套Yamantaw 	feud.Barony
}

func (c *阿什勒AshliCounty) BAshli阿什勒() feud.Barony {
	return c.阿什勒Ashli
}
    
func (c *阿什勒AshliCounty) BBlagovescen布拉戈维申() feud.Barony {
	return c.布拉戈维申Blagovescen
}
    
func (c *阿什勒AshliCounty) BBoro伯勒() feud.Barony {
	return c.伯勒Boro
}
    
func (c *阿什勒AshliCounty) BDaulakan代于莱坎() feud.Barony {
	return c.代于莱坎Daulakan
}
    
func (c *阿什勒AshliCounty) BDurtojle久尔秋利() feud.Barony {
	return c.久尔秋利Durtojle
}
    
func (c *阿什勒AshliCounty) BJanauyl亚纳乌尔() feud.Barony {
	return c.亚纳乌尔Janauyl
}
    
func (c *阿什勒AshliCounty) BTuimasy图伊马济() feud.Barony {
	return c.图伊马济Tuimasy
}
    
func (c *阿什勒AshliCounty) BYamantaw亚曼套() feud.Barony {
	return c.亚曼套Yamantaw
}
    
var CAshli阿什勒 AshliCounty = &阿什勒AshliCounty{}

func init() {
	f := CAshli阿什勒.(*阿什勒AshliCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "614",
		Title:     "ashli",
		TitleName: "阿什勒",
		TitleCode: "c_ashli",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿什勒Ashli = BAshli阿什勒
	f.阿什勒Ashli.SetParent(f)
	
	f.布拉戈维申Blagovescen = BBlagovescen布拉戈维申
	f.布拉戈维申Blagovescen.SetParent(f)
	
	f.伯勒Boro = BBoro伯勒
	f.伯勒Boro.SetParent(f)
	
	f.代于莱坎Daulakan = BDaulakan代于莱坎
	f.代于莱坎Daulakan.SetParent(f)
	
	f.久尔秋利Durtojle = BDurtojle久尔秋利
	f.久尔秋利Durtojle.SetParent(f)
	
	f.亚纳乌尔Janauyl = BJanauyl亚纳乌尔
	f.亚纳乌尔Janauyl.SetParent(f)
	
	f.图伊马济Tuimasy = BTuimasy图伊马济
	f.图伊马济Tuimasy.SetParent(f)
	
	f.亚曼套Yamantaw = BYamantaw亚曼套
	f.亚曼套Yamantaw.SetParent(f)
	
}
