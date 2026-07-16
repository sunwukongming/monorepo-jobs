package main

import (
	_ "app/init"
	"app/internal/db/mysql"
	"log"
	"sync"

	"gorm.io/gen"
)

func main() {
	db := mysql.Gorm()
	tables, err := db.Migrator().GetTables()
	if err != nil {
		log.Fatal(err)
	}
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./internal/db/mysql/query",
		Mode:          gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		ModelPkgPath:  "model",
		FieldSignable: true,
	})
	g.UseDB(db)
	var wg sync.WaitGroup
	for _, table := range tables {
		wg.Add(1)
		go func() {
			defer wg.Done()
			model := g.GenerateModel(table)
			g.ApplyBasic(model)
		}()
		wg.Wait()
	}
	g.Execute()
}
