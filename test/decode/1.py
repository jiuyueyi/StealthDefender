import sys
import re

if len(sys.argv) < 2:
    print("Usage: python script.py <payload_file>")
    sys.exit(1)

payload_file = sys.argv[1]

# 读取 Go 源文件内容
with open('.\\test\\decode\\main.go.bak', 'r', encoding='utf-8') as file:
    content = file.read()

# 读取 payload 文件内容
try:
    with open(payload_file, 'r', encoding='utf-8') as file:
        payload_text = file.read()  # 去除可能存在的前后空白字符
except FileNotFoundError:
    print(f"Error: File '{payload_file}' not found.")
    sys.exit(1)

# 构建替换字符串，确保 payload_text 中的双引号被正确转义
replacement_text = f'str_text := "{payload_text}"'
print("需要替换的字符串为:\n"+payload_text)
# 定义需要替换的变量及其新值
replacements = {
    r'str_text := ".*?"': replacement_text
}

# 查找并替换变量值
for pattern, replacement in replacements.items():
    content = re.sub(pattern, replacement, content)

# 写回更新后的内容
with open('.\\test\\decode\\script.go', 'w', encoding='utf-8') as file:
    file.write(content)

print("变量值已成功替换并更新源文件。")

