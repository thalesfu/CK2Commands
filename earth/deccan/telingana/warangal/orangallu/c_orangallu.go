package orangallu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OrangalluCounty interface {
    feud.County
    BAnumakonda阿奴摩军荼() 	feud.Barony
    BBhadrakali跋陀罗迦梨() 	feud.Barony
    BGhanpur汉普尔() 	feud.Barony
    BKothakonda拘多军荼() 	feud.Barony
    BMolangoor摩兰格尔() 	feud.Barony
    BPalampet波罗姆佩特() 	feud.Barony
    BWarangal婆浪伽罗() 	feud.Barony
}

type 乌兰伽楼OrangalluCounty struct {
	feud.BaseCounty
	阿奴摩军荼Anumakonda 	feud.Barony
	跋陀罗迦梨Bhadrakali 	feud.Barony
	汉普尔Ghanpur 	feud.Barony
	拘多军荼Kothakonda 	feud.Barony
	摩兰格尔Molangoor 	feud.Barony
	波罗姆佩特Palampet 	feud.Barony
	婆浪伽罗Warangal 	feud.Barony
}

func (c *乌兰伽楼OrangalluCounty) BAnumakonda阿奴摩军荼() feud.Barony {
	return c.阿奴摩军荼Anumakonda
}
    
func (c *乌兰伽楼OrangalluCounty) BBhadrakali跋陀罗迦梨() feud.Barony {
	return c.跋陀罗迦梨Bhadrakali
}
    
func (c *乌兰伽楼OrangalluCounty) BGhanpur汉普尔() feud.Barony {
	return c.汉普尔Ghanpur
}
    
func (c *乌兰伽楼OrangalluCounty) BKothakonda拘多军荼() feud.Barony {
	return c.拘多军荼Kothakonda
}
    
func (c *乌兰伽楼OrangalluCounty) BMolangoor摩兰格尔() feud.Barony {
	return c.摩兰格尔Molangoor
}
    
func (c *乌兰伽楼OrangalluCounty) BPalampet波罗姆佩特() feud.Barony {
	return c.波罗姆佩特Palampet
}
    
func (c *乌兰伽楼OrangalluCounty) BWarangal婆浪伽罗() feud.Barony {
	return c.婆浪伽罗Warangal
}
    
var COrangallu乌兰伽楼 OrangalluCounty = &乌兰伽楼OrangalluCounty{}

func init() {
	f := COrangallu乌兰伽楼.(*乌兰伽楼OrangalluCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1256",
		Title:     "orangallu",
		TitleName: "乌兰伽楼",
		TitleCode: "c_orangallu",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿奴摩军荼Anumakonda = BAnumakonda阿奴摩军荼
	f.阿奴摩军荼Anumakonda.SetParent(f)
	
	f.跋陀罗迦梨Bhadrakali = BBhadrakali跋陀罗迦梨
	f.跋陀罗迦梨Bhadrakali.SetParent(f)
	
	f.汉普尔Ghanpur = BGhanpur汉普尔
	f.汉普尔Ghanpur.SetParent(f)
	
	f.拘多军荼Kothakonda = BKothakonda拘多军荼
	f.拘多军荼Kothakonda.SetParent(f)
	
	f.摩兰格尔Molangoor = BMolangoor摩兰格尔
	f.摩兰格尔Molangoor.SetParent(f)
	
	f.波罗姆佩特Palampet = BPalampet波罗姆佩特
	f.波罗姆佩特Palampet.SetParent(f)
	
	f.婆浪伽罗Warangal = BWarangal婆浪伽罗
	f.婆浪伽罗Warangal.SetParent(f)
	
}
