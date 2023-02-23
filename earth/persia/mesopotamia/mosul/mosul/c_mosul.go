package mosul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MosulCounty interface {
	feud.County
	BAqrah阿奇拉() feud.Barony
	BArbil阿尔贝拉() feud.Barony
	BBakhdida巴克迪达() feud.Barony
	BBaqofah巴喀法() feud.Barony
	BBartella巴尔特拉() feud.Barony
	BKaramlish卡拉姆里斯() feud.Barony
	BMosul摩苏尔() feud.Barony
	BShekhan什汗() feud.Barony
}

type 摩苏尔MosulCounty struct {
	feud.BaseCounty
	阿奇拉Aqrah       feud.Barony
	阿尔贝拉Arbil      feud.Barony
	巴克迪达Bakhdida   feud.Barony
	巴喀法Baqofah     feud.Barony
	巴尔特拉Bartella   feud.Barony
	卡拉姆里斯Karamlish feud.Barony
	摩苏尔Mosul       feud.Barony
	什汗Shekhan      feud.Barony
}

func (c *摩苏尔MosulCounty) BAqrah阿奇拉() feud.Barony {
	return c.阿奇拉Aqrah
}

func (c *摩苏尔MosulCounty) BArbil阿尔贝拉() feud.Barony {
	return c.阿尔贝拉Arbil
}

func (c *摩苏尔MosulCounty) BBakhdida巴克迪达() feud.Barony {
	return c.巴克迪达Bakhdida
}

func (c *摩苏尔MosulCounty) BBaqofah巴喀法() feud.Barony {
	return c.巴喀法Baqofah
}

func (c *摩苏尔MosulCounty) BBartella巴尔特拉() feud.Barony {
	return c.巴尔特拉Bartella
}

func (c *摩苏尔MosulCounty) BKaramlish卡拉姆里斯() feud.Barony {
	return c.卡拉姆里斯Karamlish
}

func (c *摩苏尔MosulCounty) BMosul摩苏尔() feud.Barony {
	return c.摩苏尔Mosul
}

func (c *摩苏尔MosulCounty) BShekhan什汗() feud.Barony {
	return c.什汗Shekhan
}

var CMosul摩苏尔 MosulCounty = &摩苏尔MosulCounty{}

func init() {
	f := CMosul摩苏尔.(*摩苏尔MosulCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "697",
		Title:     "mosul",
		TitleName: "摩苏尔",
		TitleCode: "c_mosul",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿奇拉Aqrah = BAqrah阿奇拉
	f.阿奇拉Aqrah.SetParent(f)

	f.阿尔贝拉Arbil = BArbil阿尔贝拉
	f.阿尔贝拉Arbil.SetParent(f)

	f.巴克迪达Bakhdida = BBakhdida巴克迪达
	f.巴克迪达Bakhdida.SetParent(f)

	f.巴喀法Baqofah = BBaqofah巴喀法
	f.巴喀法Baqofah.SetParent(f)

	f.巴尔特拉Bartella = BBartella巴尔特拉
	f.巴尔特拉Bartella.SetParent(f)

	f.卡拉姆里斯Karamlish = BKaramlish卡拉姆里斯
	f.卡拉姆里斯Karamlish.SetParent(f)

	f.摩苏尔Mosul = BMosul摩苏尔
	f.摩苏尔Mosul.SetParent(f)

	f.什汗Shekhan = BShekhan什汗
	f.什汗Shekhan.SetParent(f)

}
