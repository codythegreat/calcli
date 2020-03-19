import subprocess

commands = [
    ['./calcli', '-answer', '3'],
    ['./calcli', '-answer', '3+3'],
    ['./calcli', '-answer', '3^{2}'],
    ['./calcli', '-answer', '2+2+2+2+2+2'],
    ['./calcli', '-answer', '5^{1}+20-10'],
    ['./calcli', '-answer', '18*0+3^{3}-9'],
]

for com in commands:
    print(subprocess.check_output(com))
