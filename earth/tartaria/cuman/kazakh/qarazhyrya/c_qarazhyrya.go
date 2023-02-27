package qarazhyrya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QarazhyryaCounty interface {
    feud.County
    BAktau_qarazhyrya阿克套() 	feud.Barony
    BGulshat_qarazhyrya古尔沙特() 	feud.Barony
    BKyzyltau克孜勒套() 	feud.Barony
    BOrtau奥尔套() 	feud.Barony
    BQarazhyrya喀拉哲拉() 	feud.Barony
    BSaryshagan萨雷沙甘() 	feud.Barony
    BZhambul江布尔() 	feud.Barony
}

type 喀拉哲拉QarazhyryaCounty struct {
	feud.BaseCounty
	阿克套Aktau_qarazhyrya 	feud.Barony
	古尔沙特Gulshat_qarazhyrya 	feud.Barony
	克孜勒套Kyzyltau 	feud.Barony
	奥尔套Ortau 	feud.Barony
	喀拉哲拉Qarazhyrya 	feud.Barony
	萨雷沙甘Saryshagan 	feud.Barony
	江布尔Zhambul 	feud.Barony
}

func (c *喀拉哲拉QarazhyryaCounty) BAktau_qarazhyrya阿克套() feud.Barony {
	return c.阿克套Aktau_qarazhyrya
}
    
func (c *喀拉哲拉QarazhyryaCounty) BGulshat_qarazhyrya古尔沙特() feud.Barony {
	return c.古尔沙特Gulshat_qarazhyrya
}
    
func (c *喀拉哲拉QarazhyryaCounty) BKyzyltau克孜勒套() feud.Barony {
	return c.克孜勒套Kyzyltau
}
    
func (c *喀拉哲拉QarazhyryaCounty) BOrtau奥尔套() feud.Barony {
	return c.奥尔套Ortau
}
    
func (c *喀拉哲拉QarazhyryaCounty) BQarazhyrya喀拉哲拉() feud.Barony {
	return c.喀拉哲拉Qarazhyrya
}
    
func (c *喀拉哲拉QarazhyryaCounty) BSaryshagan萨雷沙甘() feud.Barony {
	return c.萨雷沙甘Saryshagan
}
    
func (c *喀拉哲拉QarazhyryaCounty) BZhambul江布尔() feud.Barony {
	return c.江布尔Zhambul
}
    
var CQarazhyrya喀拉哲拉 QarazhyryaCounty = &喀拉哲拉QarazhyryaCounty{}

func init() {
	f := CQarazhyrya喀拉哲拉.(*喀拉哲拉QarazhyryaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1879",
		Title:     "qarazhyrya",
		TitleName: "喀拉哲拉",
		TitleCode: "c_qarazhyrya",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克套Aktau_qarazhyrya = BAktau_qarazhyrya阿克套
	f.阿克套Aktau_qarazhyrya.SetParent(f)
	
	f.古尔沙特Gulshat_qarazhyrya = BGulshat_qarazhyrya古尔沙特
	f.古尔沙特Gulshat_qarazhyrya.SetParent(f)
	
	f.克孜勒套Kyzyltau = BKyzyltau克孜勒套
	f.克孜勒套Kyzyltau.SetParent(f)
	
	f.奥尔套Ortau = BOrtau奥尔套
	f.奥尔套Ortau.SetParent(f)
	
	f.喀拉哲拉Qarazhyrya = BQarazhyrya喀拉哲拉
	f.喀拉哲拉Qarazhyrya.SetParent(f)
	
	f.萨雷沙甘Saryshagan = BSaryshagan萨雷沙甘
	f.萨雷沙甘Saryshagan.SetParent(f)
	
	f.江布尔Zhambul = BZhambul江布尔
	f.江布尔Zhambul.SetParent(f)
	
}
