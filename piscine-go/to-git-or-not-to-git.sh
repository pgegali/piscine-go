#!/bin/bash

# Задайте URL-адрес JSON-файла
json_url="https://zero.academie.one/assets/superhero/all.json"

# Используйте curl для загрузки JSON-файла и jq для извлечения данных
curl -s "$json_url" | jq -r '.[] | select(.id == 170) | .name, .power, .gender'

