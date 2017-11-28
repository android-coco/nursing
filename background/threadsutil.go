package background

import (
	"fit"
	"time"
	"fmt"
)

// 员工表
type BCE1 struct {
	BCE01 uint64    //ID
	BCE02 string    //编码, 工号, Code
	BCE03 string    //姓名, 在不同场合的名称： EMPNN=护士, EMPPN=医师, EMPDN=药剂师, EMPWN=仓管员, EMPTN=医技师,
	BCE04 string    //英文名, NameB
	BCE05 string    //称谓, Title
	BCE06 string    //曾用名, PreviousName
	ABBRP string    //拼音, PYImc
	ABBRW string    //五笔, WBImc
	ABW01 string    //性别, Sex
	ACK01 string    //婚姻, Marriage, Marriaged
	BCE11 string    //身份证号, IDNumber
	ABQ01 string    //民族, Nation
	BCE13 time.Time //出生日期, BirthDate
	BCK01 string    //部门ID
	ACP01 string    //政治面貌, TPLAC.PLACC, PoliticalAffiliationID
	AAY01 string    //学历, 最高学历, 参见表, TEDUL
	ACT01 string    //学位, TDEGR
	BCE18 string    //毕业学校, graduated school
	BCE19 time.Time //毕业时间, graduated date
	ABU02 string    //ADIVN
	AAD01 string    //所学专业
	ABG01 string    //从事专业, TPPST
	BCE23 string    //执业证号
	BCE24 string    //户口所在地, RegisteredResidence
	ABS01 string    //行政职务
	ABI01 string    //技术职务, 职称
	AAH01 string    //聘任职务
	BCE28 time.Time //工作日期, joindate
	AAQ01 string    //执业类别
	ABE01 string    //执业范围, 需要单独的表
	BCE31 time.Time //进院日期, HireDate
	BCE32 time.Time //离职日期, DimissionDate
	BCE33 string    //离职说明, DimissionDescription
	BCE34 string    //住址, Adress
	BCE35 string    //办公电话, officephone
	BCE36 string    //联系电话, phone
	BCE37 string    //移动电话, mobilephone
	BCE38 string    //电子邮箱, email
	BCE39 string    //备注, comment
	BCE41 int64     // --状态, 0=试用,1=在职, 2=离职, 3=退休
}

// 病人表
type VAA1 struct {
	VAA01  uint64    // 病人ID, PATFU, PATID
	VAA02  string    // 会员卡号
	VAA03  string    // 门诊号 Clinic Patient No, Out patient No, OPTNO
	VAA04  string    // 住院号, In Patient No, IPTNO
	VAA05  string    // 姓名, PATNM, FirstName + Middlename
	VAA06  string    // 监护人1
	ABBRP  string    // 拼音
	ABBRW  string    // 五笔
	ABW01  string    // 性别
	VAA10  int       // 年龄
	AAU01  string    // 年龄单位
	VAA12  time.Time // 出生日期, birth date
	ACK01  string    // 婚姻, marital status
	VAA14  string    // 身份证件, 指公安机关签发的有效证件
	VAA15  string    // 身份证号, IDNo, IDNumber
	VAA16  string    // 其他证件
	ABJ01  string    // 付款方式, 医疗付款方式
	BDP02  string    // 病人类别
	ABC02  string    // 病人费别
	VAA20  string    // 出生地点, 出生地, BirthPlace
	ACM02  string    // 从业状况, 身份
	ACC02  string    // 国籍, Nationality
	ABQ02  string    // 民族, Nation
	VAA25  string    // 籍贯, 省.市, NativePlace
	VAA26  string    // 宗教, 佛教, 伊斯兰教, 基督教, 天主教, 犹太教, 耶稣教, 其他, Religion, 暂时不用
	VAA27  string    // 种族, 人种, Ethnic, 暂时不用
	VAA28  string    // 户口地址, 户籍地址, RegisteredAddress
	VAA29  string    // 户籍邮编, RegisteredPostCode
	VAA30  string    // 户籍电话, RegisteredPhone
	VAA31  string    // 省市
	VAA32  string    // 县市, 城市, 县区市
	VAA33  string    // 地址, 常住地址, resident address
	VAA34  string    // 电话
	VAA35  string    // 移动电话
	VAA36  string    // 电子邮箱
	VAA37  string    // 其他联系方式, 如QQ号, MSN号, Skype号等, OtherCommunications
	VAA38  string    // 学历
	VAA39  string    // 监护人2, Guardian
	VAA40  string    // 联系人姓名, contact person name
	VAA41  string    // 与病人关系, RelationShip
	VAA42  string    // 联系人地址, Contact person address
	VAA43  string    // 联系人电话, contact person telephone
	VAA44  string    // 联系人移动电话, Contact person Mobile Phone
	BAQ01  int       // 合同单位ID
	BAQ02  string    // 单位编码, CoCode
	VAA47  string    // 工作单位, CoName
	VAA48  string    // 单位电话, CoTelephone
	VAA49  string    // 单位地址, CoAddress
	VAA50  string    // 单位邮编, CoPostCode
	VAA51  string    // 单位开户行
	VAA52  string    // 单位银行帐号
	VAA53  string    // 担保人, Guarantor
	VAA54  float32   // 信用额度, 担保额度, CreditLimit
	VAA55  int       // 担保性质, CreditType
	VAA56  int       // 住院次数, HospitalizationNumber
	VAA57  time.Time // 就诊时间, LastVisitDate
	BCK01A int       // 就诊科室, LastVisitDeptID
	VAA61  int       // 就诊状态, 0=无, 1=门诊, 2=住院, 3=出院, 4=转院, 5=死亡, 9=其他 VisitState
	VAA62  string    // 过敏史
	BDX02  string    // 了解途径, 病人了解医院的方式(如电视广告, 介绍, 户外广告等)
	VAA64  string    // 备注
	VBU01  int       // 帐号ID
	VAA66  string    // 病案号
	VAA67  string    // 查询密码
	IAK05  string    // 社会保障号
	IAA01  int       // 保险机构
	BCK01B int       // 科室ID
	BCK01C int       // 病区ID
	BCQ04  string    // 床号
	VAA73  time.Time // 入院时间
	VAA74  time.Time // 出院时间
	VAA75  time.Time // 建档时间
	BEP05  float32   // 住院报警值
	BEP06  float32   // 住院信用额度
	ABL01A string    // 正定型
	ABL01B string    // 反定型
	VAA76  time.Time // 有效时间
	ABL01  string    // 血型
	BEP06B float32   // 门诊信用额度
	VAA78  string    // Rh血型
	VAA82  string    // 健康卡号
	VAA01A int       // 相关ID
	VAA84  string    // 体检登记号
}

//
//type User struct {
//	Username   string
//	Code       string
//	Password   string
//	Employeeid uint64
//	Authority  int
//}

func init() {
	fit.RegisterMime()
	// 请不要随意创建新的线程
	go synchronizing()
}

//同步用户信息到我们系统
func synchronizing() {
	for {
		if !fit.SartOK {
			continue
		}
		synchronizingBCE1()
		synchronizingVAA1()
		time.Sleep(time.Minute * time.Duration(fit.Config().Cycle)) // 停顿5分钟
	}
}

func synchronizingBCE1() {
	mods := make([]BCE1, 0)
	timeMark := time.Now()
	fit.SQLServerEngine().SQL("SELECT * FROM BCE1").Find(&mods)
	for _, k := range mods {
		_, err := fit.MySqlEngine().Exec("delete from BCE1 where BCE01 = ?", k.BCE01)
		_, err = fit.MySqlEngine().Table("BCE1").InsertOne(&k)
		if err != nil {
			fmt.Println("***JK", err.Error())
		}
	}
	timeNow := time.Now()
	timeDur := timeNow.Sub(timeMark)

	str := fmt.Sprintf("耗时：%.1f秒", timeDur.Seconds())
	fit.Logger().LogError("BackgorundThread：", "同步BCE1结束", "数据条数：", len(mods), str)
}

/*同步病人信息到我们系统*/
func synchronizingVAA1() {
	mods := make([]VAA1, 0)
	timeMark := time.Now()

	// 查询所有病区在院+入院的病人
	fit.SQLServerEngine().SQL("select VAA1.* from BCK1, VAE1, VAA1 where BCK1.BCK01A = 141 and VAE1.BCK01C = BCK1.BCK01 and VAE44 in ('1','2') and VAA1.VAA01 = VAE1.VAA01").Find(&mods)
	//fit.SQLServerEngine().SQL("SELECT * FROM VAA1").Find(&mods)
	for _, k := range mods {
		_, err := fit.MySqlEngine().Exec("delete from VAA1 where VAA01 = ?", k.VAA01)
		_, err = fit.MySqlEngine().Table("VAA1").InsertOne(&k)
		if err != nil {
			fmt.Println("***JK", err.Error(), "  ", k)
		}
	}
	timeNow := time.Now()
	timeDur := timeNow.Sub(timeMark)

	str := fmt.Sprintf("耗时：%.1f秒", timeDur.Seconds())
	fit.Logger().LogError("BackgorundThread：", "同步VAA1结束", "数据条数：", len(mods), str)
}
