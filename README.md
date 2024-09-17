## Подключение
    package main

    import (
        "fmt"

        "github.com/artlevitan/go-helpers" // Импорт пакета
    )

    func main() {
        // Пример использования функции для работы с хэшами
        hash := helpers.MD5("example")
        fmt.Println("MD5 hash:", hash) // 1a79a4d60de6718e8e5b326e338ae533

        // Пример использования функции для валидации
        isValidURL := helpers.IsURL("https://www.appercase.ru")
        fmt.Println("Is valid URL:", isValidURL) // true

        // Пример использования функции для работы с числами
        rounded := helpers.RoundFloat(123.456, 2)
        fmt.Println("Rounded float:", rounded) // 123.46

        // Пример использования функции для работы со строками
        cutString := helpers.CutString("This is a long string", 10)
        fmt.Println("Cut string:", cutString) // This is a
    }

## Лицензия
Copyright 2023-2024, Appercase LLC. All rights reserved.
https://www.appercase.ru/