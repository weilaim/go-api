package model

//执行自动迁移

func migration(){
	//自动迁移模式
	_ = DB.AutoMigrate(&User{})
	//自动迁移Video表
	_ = DB.AutoMigrate(&Video{})
}