package scripts

//func main() {
//	db := GetDatabaseConnection()
// 建表
//db.AutoMigrate(new(etity.User))
//db.AutoMigrate(&etity.Lesson{})

//增加
//db.Create(&[]etity.User{
//	{Username: "哈哈"},
//	{Username: "呵呵4"},
//	{Username: "嘻嘻"},
//	{Username: "啦啦"},
//	{Username: "嚯嚯"},
//})
//s := etity.Student{
//	Model:         gorm.Model{
//		ID: 1,
//	},
//	StudentName:   "琪琪",
//	StudentAge:    18,
//	StudentEmail:  "qq@example.com",
//	StudentMobile: "7003823823",
//}
//
//l := etity.Lesson{
//	Model:         gorm.Model{
//		ID : 1,
//	},
//	Student: s,
//}
//db.Create(&l)

// 查询
//var user etity.User
//var results map[string]interface{}
// 主键检索
//dbres := db.Model(&etity.User{}).First(&results)
// 条件检索(string)
//db.Where("username = ? AND id = ?","哈哈",3).First(&user)
//fmt.Println(user)
//fmt.Print(results)
//fmt.Println(errors.Is(dbres.Error,gorm.ErrRecordNotFound))
// 自定义结构体查询 要加model方法

//更新
// update 只更新选择的字段
// updates 更新所有 一种为map 一种为结构体
// save 无论什么都更新
//db.Model(&etity.User{}).Where("username = ? ","哈哈").Update("username","芜湖")

//删除
//db.Delete(&etity.User{},2)

//}
