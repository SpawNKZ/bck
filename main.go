package main

import (
	"fmt"
	"sort"
	"strings"
)

type FileStruct struct {
	Size uint
	Name string
}

func makeUniqueFileNames(files []FileStruct) {
	// Сортируем слайс по размеру файла в порядке убывания
	sort.SliceStable(files, func(i, j int) bool {
		return files[i].Size > files[j].Size
	})

	// Создаем мапу для отслеживания уже использованных имен файлов
	usedNames := make(map[string]bool)

	// Проходим по каждому файлу и делаем его имя уникальным
	for i := 0; i < len(files); i++ {
		file := &files[i]

		// Проверяем, является ли текущее имя файла уникальным
		_, isUsed := usedNames[file.Name]

		// Если имя уже использовано или содержит суффикс " (", добавляем суффикс " (%n)"
		if isUsed || strings.Contains(file.Name, " (") {
			suffix := 3 // Начинаем с суффикса " (3)"
			newName := file.Name

			// Ищем уникальное имя, пока не найдем свободное
			for usedNames[newName] {
				newName = fmt.Sprintf("%s (%d)", file.Name, suffix)
				suffix++
			}

			// Обновляем имя файла
			file.Name = newName
		}

		// Отмечаем текущее имя файла как использованное
		usedNames[file.Name] = true
	}
}
