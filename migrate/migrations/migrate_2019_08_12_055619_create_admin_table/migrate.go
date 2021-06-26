package migrate_2019_08_12_055619_create_admin_table

import "gin-web/migrate"

func Up() {

	migrate.Create("admin", func(createMigrate *migrate.Migrate) {

		createMigrate.Name = "migrate_2019_08_12_055619_create_admin_table"

		createMigrate.BigIncrements("id")

		createMigrate.Integer("user_id").Unsigned().Nullable().Default(0).Comment("用户id")

	})

}

func Down() {

	migrate.DropIfExists("admin")

}
