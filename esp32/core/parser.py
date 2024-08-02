import requests
import re

# URL директории с исходным кодом
base_url = "https://raw.githubusercontent.com/espressif/esptool/master/esptool/targets/"

# Получаем содержимое файла __init__.py
init_url = base_url + "__init__.py"
response = requests.get(init_url)
init_content = response.text

# Находим все строки импорта
import_lines = [line.strip() for line in init_content.splitlines() if line.startswith("from .")]

# Извлекаем названия модулей, удаляя точку и добавляя расширение .py
module_names = [line.split()[1][1:] + ".py" for line in import_lines]

# Скачиваем и анализируем каждый модуль
for module_name in module_names:
    module_url = base_url + module_name
    module_response = requests.get(module_url)
    module_content = module_response.text

    # Находим определения классов в модуле
    class_names = re.findall(r"class (\w+)", module_content)

    # Для каждого класса находим и печатаем его переменные
    for class_name in class_names:
        print(f"\nПеременные из класса {class_name} в модуле {module_name}:")
        class_pattern = rf"class {class_name}\s*:\s*(.*?)\n\n"
        class_body = re.search(class_pattern, module_content, re.DOTALL)
        if class_body:
            # Извлекаем переменные класса
            variables = re.findall(r"^\s*(\w+)\s*=", class_body.group(1), re.MULTILINE)
            for variable in variables:
                variable_value = re.search(rf"^\s*{variable}\s*=\s*(.*)", class_body.group(1), re.MULTILINE)
                if variable_value:
                    print(f"{variable} = {variable_value.group(1).strip()}")
