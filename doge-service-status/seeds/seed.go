package seeds

import (
	"doge-service-status/datasource"
	"github.com/doge-soft/dogego_module_dbseeder"
)

func SeedDatabase() {
	seeder := dogego_module_dbseeder.NewSeeder(
		datasource.MasterDatabase(),
		datasource.SlaveDatabase())
}
