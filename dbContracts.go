package main

import (
	"context"
	"log"
)

// insertConfigDb добавляет если сервис существует версия плюс 1
func (d *Data) insertConfigDb(item *StructForJson) {

	lastVersion := d.getMaxVersion(item.Service)
	_, err := d.conn.Exec(context.Background(), insertConfigs, item.Service, item.Data, lastVersion+1)
	if err != nil {
		log.Fatal("error wit conn.Exec ", err)
	}
}

// getMaxVersion получает максимальную версию конфига
func (d *Data) getMaxVersion(itemName string) int {
	rows, err := d.conn.Query(context.Background(), selectMaxConfig, itemName)
	if err != nil {
		log.Println("error wit conn.Query ", err)
	}
	defer rows.Close()
	var s StructForJson
	var version int
	rows.Next()
	err = rows.Scan(&s.Service, &s.Data, &version)
	if err != nil {
		log.Println("error with rows.Scan", err)
	}

	return version
}

// selectConfig возвращает указатель на конфиг
func (d *Data) selectConfig(itemName string) *StructForJson {
	rows, err := d.conn.Query(context.Background(), selectMaxConfig, itemName)
	if err != nil {
		log.Println("error wit conn.Query ", err)
	}
	defer rows.Close()
	var s StructForJson
	var version int
	rows.Next()
	err = rows.Scan(&s.Service, &s.Data, &version)
	if err != nil {
		log.Println("error with rows.Scan", err)
	}
	return &s
}

// checkItem Удаление максимальной версии конфига
func (d *Data) deleteConfigDb(itemName string) {
	lastVersion := d.getMaxVersion(itemName)
	_, err := d.conn.Exec(context.Background(), deleteMaxConfig, itemName, lastVersion)
	if err != nil {
		log.Fatal("error wit conn.Exec ", err)
	}
}
