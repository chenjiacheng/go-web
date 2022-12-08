package seeders

import (
	"fmt"
	"go-web/app/models/user"
	"go-web/database/factories"
	"go-web/pkg/console"
	"go-web/pkg/logger"
	"go-web/pkg/seed"
	"gorm.io/gorm"
)

func init() {

	// 添加 Seeder
	seed.Add("SeedUsersTable", func(db *gorm.DB) {

		// 创建 10 个用户对象
		users := factories.MakeUsers(10)

		// 这是第 11 个，我们定制手机号和邮箱，方便测试
		userModel := user.User{
			Name:     "Summer",
			Email:    "summer@example.com",
			Phone:    "18600000000",
			Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe", // secret
		}
		users = append(users, userModel)

		// 批量创建用户（注意批量创建不会调用模型钩子）
		result := db.Table("users").Create(&users)

		// 记录错误
		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		// 打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}
